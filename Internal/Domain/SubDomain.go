package Domain

import (
	"bytes"
	"context"
	"fmt"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
	"go_test/Internal/Infomation"
	"go_test/comand"
	"go_test/logger"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type SubDomainRes struct {
	SubDomain string
	IPs       []string
}

func showLoading(stopchan chan bool) {
	frames := []string{"|", "/", "-", "\\"}
	i := 0
	for {
		select {
		case <-stopchan:
			fmt.Printf("/r")
			return

		default:
			fmt.Printf("\r[%s] Scanning...", frames[i])
			time.Sleep(100 * time.Millisecond)
			i = (i + 1) % len(frames)
		}

	}
}

func Passives(domain string) error {
	stopchan := make(chan bool)
	go showLoading(stopchan)
	subfinderOpts := &runner.Options{
		Threads:            10, // Thread controls the number of threads to use for active enumerations
		Timeout:            30, // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 10, // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		// ResultCallback: func(s *resolve.HostEntry) {
		// callback function executed after each unique subdomain is found
		// },
		// ProviderConfig: "your_provider_config.yaml",
		// and other config related options
	}

	subfinder, err := runner.NewRunner(subfinderOpts)
	if err != nil {

		logger.Log.ErrorMsgf("failed to create subfinder runner: %v", err)
	}

	output := &bytes.Buffer{}
	var sourceMap map[string]map[string]struct{}
	// To run subdomain enumeration on a single domain
	if sourceMap, err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), domain, []io.Writer{output}); err != nil {
		//log.Fatalf("failed to enumerate single domain: %v", err)
		logger.Log.ErrorMsgf("failed to enumerate single domain: %v", err)
	}

	for subdomain, sources := range sourceMap {
		sourcesList := make([]string, 0, len(sources))
		for source := range sources {
			sourcesList = append(sourcesList, source)
		}
		log.Printf("%s %s (%d)\n", subdomain, sourcesList, len(sources))
	}
	return nil
}

func Actives(o *Infomation.Options, domain string) {
	logger.Log.InfoMsgf("开启主动扫描")
	var wordlist []string
	var err error
	if o.ScanDomain.WordListPath == "" {
		WordlistUrl := "https://raw.githubusercontent.com/danielmiessler/SecLists/refs/heads/master/Discovery/DNS/subdomains-top1million-110000.txt"
		wordlist, err = comand.UrlWordlist(WordlistUrl)
		if err != nil {
			logger.Log.ErrorMsgf("failed to fetch wordlist: %v", err)
		}
	} else {
		wordlist, err = comand.LoadWordList(o.ScanDomain.WordListPath)
		if err != nil {
			logger.Log.ErrorMsgf("failed to load wordlist: %v", err)
		}
		logger.Log.InfoMsgf("加载字典成功 %s 数量为%d", o.ScanDomain.WordListPath, len(wordlist))

	}
	var results []SubDomainRes
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var scan func(string)
	scan = func(target string) {
		for _, word := range wordlist {
			subdomain := word + "." + target
			wg.Add(1)
			go func(subdomain string) {
				defer wg.Done()
				var address []string
				var err error
				address, err = net.LookupHost(subdomain)
				if err == nil {
					mutex.Lock()
					result := SubDomainRes{SubDomain: subdomain}
					if o.ScanDomain.ShowIp {
						result.IPs = address
					}
					results = append(results, result)
					mutex.Unlock()
					if o.ScanDomain.ShowIp {
						logger.Log.InfoMsgf("子域名:%s,IP:%s", subdomain, address)
					} else {
						logger.Log.InfoMsgf("子域名:%s", subdomain)
					}
				}

			}(subdomain)
		}
	}
	scan(domain)
	wg.Wait()
}
