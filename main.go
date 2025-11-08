package main

import (
	"github.com/jessevdk/go-flags"
	"go_test/Internal/Domain"
	"go_test/Internal/Infomation"
	"go_test/comand"
	"go_test/logger"
	"log/slog"
	"os"
)

func main() {
	logger.Init(logger.WithLevel(slog.LevelDebug),
		logger.WithTimeFormat("15:04:05"),
		logger.WithOutPutJson(true),
		logger.WithUserColor(true),
	)
	//logger.Log.Debug("Debug message", slog.String("key", "value"))
	//logger.Log.DebugMsgf("Debug message %s", "formatted")
	var opt Infomation.Options
	opt.Cfg = *comand.Parse()
	parser := flags.NewParser(&opt, flags.Default)
	parser.Parse()

	if len(os.Args) == 1 {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}
	if len(opt.Info.Icp) > 0 {
		slice := Infomation.TraverseSlice(opt.Info.Icp)
		for _, s := range slice {
			Infomation.QueryDomain(s, &opt)
		}
	}
	if len(opt.Info.Ip) > 0 {
		slice := Infomation.TraverseSlice(opt.Info.Ip)
		for _, s := range slice {
			Infomation.QueryIp(&opt, s)
		}
	}
	if len(opt.Info.Whois) > 0 {
		slice := Infomation.TraverseSlice(opt.Info.Whois)
		for _, s := range slice {
			Infomation.Whois(s)
		}
	}
	if opt.ScanDomain.Active && opt.ScanDomain.Passive {
		logger.Log.ErrorMsgf("不能同时使用主动和被动扫描")
	} else if opt.ScanDomain.Passive {
		if len(opt.ScanDomain.Domain) > 0 {
			slice := Infomation.TraverseSlice(opt.ScanDomain.Domain)
			for _, s := range slice {
				Domain.Passives(s)
			}
		}

	} else if opt.ScanDomain.Active {
		if len(opt.ScanDomain.Domain) > 0 {
			slice := Infomation.TraverseSlice(opt.ScanDomain.Domain)
			for _, s := range slice {
				Domain.Actives(&opt, s)
			}
		}
	}

}
