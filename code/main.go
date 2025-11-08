package main

import (
	"code/dirsearch"
	"code/ps"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"sync"
)

type Ip struct {
	Ip        string `short:"i" long:"ip" description:"输入扫描IP"`
	Startport int    `short:"s" long:"startport" description:"开始端口" default:"1" `
	Endport   int    `short:"e" long:"endport" description:"结束端口" default:"65535" `
}
type Ipinfo struct {
	GetIp Ip             `command:"portscan" description:"端口扫描"`
	Dir   dirsearch.Dirs `command:"dir" description:"目录遍历"`
}

func (a *Ipinfo) Execute(args []string) error {

	if a.GetIp.Ip != "" {
		var wg sync.WaitGroup
		for i := a.GetIp.Startport; i <= a.GetIp.Endport; i++ {
			wg.Add(1)
			go ps.ScanPort(a.GetIp.Ip, i, &wg)
		}
		wg.Wait()
		fmt.Println("扫描结束")
	}
	if a.Dir.Path != "" {
		dirsearch.Dirsearch(&a.Dir)
	}

	return nil
}

func main() {
	var opt Ipinfo
	parser := flags.NewParser(&opt, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
	}
	opt.Execute(os.Args[1:])
}
