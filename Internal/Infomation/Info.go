package Infomation

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
	"go_test/logger"
	"strings"
)

var (
	a1 string
	a2 string
	a3 string
	a4 string
	a5 string
	a6 string
	a7 string
)

func QueryDomain(domin string, o *Options) {
	client := req.C()
	//client.SetProxyURL("http://127.0.0.1:1111")
	if o.Req.Proxy != "" {
		client.SetProxyURL(o.Req.Proxy) //设置代理
	}
	request := client.R()
	url := "https://www.beianx.cn/search/" + domin
	fmt.Println(url)
	get, err := request.SetHeaders(map[string]string{
		"Cookie": fmt.Sprintf(" machine_str=%s", o.Cfg.Cookies.BeiAn),
	}).Get(url) //常规的发请求 爬虫

	if err != nil {
		logger.Log.ErrorMsgf("请求失败 %s", err)
	}
	if !get.IsSuccessState() {
		logger.Log.ErrorMsgf("请求失败 状态码是:%s", get.Status)
	}
	if get.StatusCode == 401 {
		logger.Log.ErrorMsgf("Cookies失效 请重新获取")
	}
	reader, err := goquery.NewDocumentFromReader(get.Body)
	if err != nil {
		logger.Log.ErrorMsgf("HTML解析错误：%s", err)
	}

	reader.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		a1 = strings.TrimSpace(s.Find("td").Eq(1).Text())
		a2 = strings.TrimSpace(s.Find("td").Eq(2).Text())
		a3 = strings.TrimSpace(s.Find("td").Eq(3).Text())
		a4 = strings.TrimSpace(s.Find("td").Eq(4).Text())
		a5 = strings.TrimSpace(s.Find("td").Eq(5).Text())
		a6 = strings.TrimSpace(s.Find("td").Eq(6).Text())
		a7 = strings.TrimSpace(s.Find("td").Eq(7).Text())

	})

	logger.Log.InfoMsgf(
		"+---------------------------------------------\n"+
			"                 |主办单位名称   ｜%s       \n"+
			"                 |主办单位性质   ｜%s       \n"+
			"                 |网站备案号     ｜%s       \n"+
			"                 |网站名称       ｜%s       \n"+
			"                 |网站首页地址   ｜%s       \n"+
			"                 |审核日期       ｜%s       \n"+
			"                 |是否限制接入   ｜%s       \n"+
			"                 +---------------------------------------------", a1, a2, a4, a5, a6, a7)
	o.Res.IcpRes.OrganizationName = a1
	o.Res.IcpRes.OrganizationType = a2
	o.Res.IcpRes.IcpNumber = a3
	o.Res.IcpRes.WebSiteName = a4
	o.Res.IcpRes.HomePageUrl = a5
	o.Res.IcpRes.ReviewDate = a6
	o.Res.IcpRes.AccessRestriction = a7

}
