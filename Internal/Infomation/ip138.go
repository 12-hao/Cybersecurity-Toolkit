package Infomation

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
	"go_test/logger"
	"strings"
)

func QueryIp(o *Options, ip string) error {
	client := req.C()
	//client.SetProxyURL("http://127.0.0.1:1111")
	if o.Req.Proxy != "" {
		client.SetProxyURL(o.Req.Proxy)
	}
	request := client.R()
	get, err := request.Get("https://site.ip138.com/" + ip)
	if err != nil {
		return logger.Log.ErrorMsgf("网络请求失败", err)
	}
	reader, err := goquery.NewDocumentFromReader(get.Body)
	if err != nil {
		return logger.Log.ErrorMsgf("解析HTML失败", err)
	}
	text := reader.Find("h3").First().Text()
	o.Res.IpRes.Address = text
	noResult := reader.Find("#list li").Last().Text()
	if noResult == "暂无结果" {
		o.Res.IpRes.Address = strings.TrimSpace(o.Res.IpRes.Address)
		logger.Log.InfoMsgf(
			"IP %s 查询结果\n"+
				"+-------------+------------------\n"+
				"|IP地址        |%s\n"+
				"|归属地        |%s\n"+
				"+-------------+------------------\n"+
				"未查询到相关绑定信息\n", ip, ip, o.Res.IpRes.Address)
	} else {
		var bindtimes []string
		var bindsites []string
		o.Res.IpRes.Address = strings.TrimSpace(o.Res.IpRes.Address)
		reader.Find("#list li").Each(func(i int, s *goquery.Selection) {
			if i < 2 {
				return
			}
			date := s.Find(".date").Text()
			site := s.Find("a").Text()
			if date != "" && site != "" {
				bindtimes = append(bindtimes, strings.TrimSpace(date))
				bindsites = append(bindsites, site)
			}

		})
		o.Res.IpRes.BindTimes = bindtimes
		o.Res.IpRes.BindSites = bindsites
		logger.Log.InfoMsgf("绑定时间", bindtimes)
		logger.Log.InfoMsgf("绑定网站", bindsites)
		logger.Log.InfoMsgf("共有%d个绑定 具体信息如下", len(bindtimes))
		for i2, _ := range bindtimes {

			logger.Log.InfoMsgf(
				"+--------------------------------\n"+
					"                 |绑定时间   |%s\n"+
					"                 |绑定的网站｜%s\n"+
					"                 +--------------------------------\n", bindtimes[i2], bindsites[i2])
		}
	}
	return nil
}
