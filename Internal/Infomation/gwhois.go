package Infomation

import (
	"fmt"
	"github.com/likexian/whois"
	"go_test/logger"
)

func Whois(domain string) {
	result, err := whois.Whois(domain)
	if err != nil {
		logger.Log.ErrorMsgf("无法找到数据", err)
	}
	fmt.Println(result)
}
