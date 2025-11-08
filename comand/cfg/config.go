package cfg

type Config struct {
	Cookies struct {
		BeiAn string `yaml:"BeiAn"`
	} `yaml:"Cookies"`
} //定义一个像json一样的读取文件方式

var CfgName = "config.yaml" //定义文件的名字

var CfgYaml = `Cookies:
  BeiAn: ""` //定义文件的的格式后直接读取

var Version = "1.0.0" //定义版本信息
