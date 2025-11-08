package comand

import (
	"bufio"
	"github.com/imroc/req/v3"
	"go_test/comand/cfg"
	"go_test/comand/utils"
	"go_test/logger"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"strings"
)

func Parse() *cfg.Config { //调用函数读取配置文件 调用的是cfg里面的Config这个结构体
	exists, _ := utils.PathExists(cfg.CfgName) //判断文件是否存在
	if !exists {
		file, err := os.Create(cfg.CfgName)        //如果文件不存在就重新创建
		_, err = io.WriteString(file, cfg.CfgYaml) //并且写入内容
		if err != nil {
			logger.Log.Error("创建配置文件失败%s", err)
		}
	}
	readFile, err := os.ReadFile(cfg.CfgName) //如果存在就直接读取这个文件的名字 是上面文件里面定义的文件名字
	if err != nil {
		logger.Log.Error("读取文件失败%s", err)
	}
	c := new(cfg.Config) //new一个结构体 里面存放cookie和BeiAn
	if err := yaml.Unmarshal(readFile, c); err != nil {
		logger.Log.ErrorMsgf("解析配置文件失败%s", err)
	}
	return c
}

func UrlWordlist(url string) ([]string, error) {
	client := req.C()
	r := client.R()
	res, err := r.Get(url)
	if err != nil {
		logger.Log.ErrorMsgf("请求字典失败", err)
	}
	var words []string
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			words = append(words, line)
		}
	}
	return words, nil
}

func LoadWordList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		logger.Log.ErrorMsgf("打开字典失败", err)
	}
	defer file.Close()
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			words = append(words, line)
		}
	}
	return words, nil
}
