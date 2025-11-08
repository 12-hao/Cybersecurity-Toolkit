package dirsearch

import (
	"bufio"
	"fmt"
	"github.com/imroc/req/v3"
	"os"
	"strings"
)

type Dirs struct {
	Path    string `short:"p" long:"paht" description:"设置字典路径"`
	Address string `short:"a" long:"address" description:"设置要爆破的url"`
}

func Dirsearch(d *Dirs) {
	var path []string
	file, err := os.OpenFile(d.Path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) //读取文件
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "/") { //判断是否是/开头 影响拼接
			line = strings.TrimPrefix(line, "/")
		}
		path = append(path, line)
	}
	if strings.HasSuffix(d.Address, "/") { //判断是否是/开头 影响拼接
		d.Address = strings.TrimSuffix(d.Address, "/")
	}
	for _, u := range path {
		url := d.Address + "/" + u
		client := req.C()
		r := client.R()
		get, err := r.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("正在访问 ", url)
		if get.StatusCode == 200 || get.StatusCode == 403 || get.StatusCode == 405 {
			fmt.Printf("%s该路径存在 访问状态码是%d \n", url, get.StatusCode)
		}
	}

}

//func main() {
//	dirs := Dirs{
//		Path:    "1.txt",
//		Address: "http://www.bing.com",
//	}
//	Dirsearch(&dirs)
//}
