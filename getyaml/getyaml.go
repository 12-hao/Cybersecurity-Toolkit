package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	HttpMatcher []struct {
		Raw []string `yaml:"raw"`
	} `yaml:"http"`
}

func main() {
	file, err := os.ReadFile("CVE-2024-38856.yaml")
	if err != nil {
		fmt.Println("文件打开错误", err)
	}
	/*scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}*/
	h := new(Config)
	err = yaml.Unmarshal(file, &h)
	if err != nil {
		fmt.Println("解析yaml文件失败", err)
	}
	fmt.Println(h.HttpMatcher[0].Raw[0])
}
