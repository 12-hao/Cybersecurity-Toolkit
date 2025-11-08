package Infomation

import (
	"go_test/comand/cfg"
	"go_test/logger"
	"reflect"
)

type IPResult struct {
	Address   string
	BindTimes []string
	BindSites []string //接受结果的结构体 在下面集合了
}

type IcpResult struct {
	OrganizationName  string
	OrganizationType  string
	IcpNumber         string
	WebSiteName       string
	HomePageUrl       string
	ReviewDate        string
	AccessRestriction string
}

type Result struct {
	IpRes  IPResult
	IcpRes IcpResult
}

type SubDomain struct {
	Passive      bool     `short:"p" long:"passive" description:"被动获取子域名"`
	Active       bool     `short:"a" long:"active" description:"主动获取子域名"`
	Domain       []string `short:"d" long:"domain" description:"输入域名"`
	WordListPath string   `short:"w" long:"wordlistpath" description:"字典目录"`
	ShowIp       bool     `short:"s" long:"showip" description:"是否显示IP"`
}

type RequestCommand struct {
	Proxy string `short:"p" long:"proxy" description:"设置代理 要带上协议头"`
}

type InFoCommand struct {
	Icp   []string `short:"i" long:"icp" description:"查询域名备案信息"`     //参数-i
	Ip    []string `short:"n" long:"ip" description:"查询ip绑定域名"`      //参数-n
	Whois []string `short:"w" long:"whois" description:"查询域名以及IP信息"` //参数-w
}
type ConfigCommand struct {
	Version bool `short:"v" long:"version" description:"查看版本信息"`
}

type Options struct {
	Info       InFoCommand   `command:"info" description:"查询信息"`
	Config     ConfigCommand `command:"config" description:"配置信息"`
	Cfg        cfg.Config    //定义一个结构体来存储cookies
	Res        Result        //用来接受结果
	Req        RequestCommand
	ScanDomain SubDomain `command:"subdomain" description:"查询子域名"`
}

func TraverseSlice(slice interface{}) []string { //定义一个任意类型的的函数 返回一个切片
	var result []string
	v := reflect.ValueOf(slice)                   //这个函数是可以将slice里面的值转换到V上去
	if v.Kind() == reflect.Slice && v.Len() > 0 { //判断值的类型是否一样 并且切片长度是否不是0
		for i := 0; i < v.Len(); i++ {
			result = append(result, v.Index(i).String())
		}
	}
	return result
}

/*func (i *InFoCommand) Execute(args []string) error {
	i.cfg = *comand.Parse()            //执行命令的时候去获取这个cfg
	traversal := TraverseSlice(i.Icp)  //利用函数去遍历出来的切片
	for _, domina := range traversal { //切片里的每一个值都查询
		icp(domina, i)
	}
	Ip := TraverseSlice(i.Ip)
	if len(i.Ip) > 0 {
		for _, ip := range Ip {
			QueryIp(i, ip)
		}
	}
	Whoi := TraverseSlice(i.Whois)
	if len(i.Whois) > 0 {
		for _, w := range Whoi {
			Whois(w)
		}
	}
	return nil

}*/

func (c *ConfigCommand) Execute(args []string) error { //执行version这个命令输出版本号
	if c.Version {
		logger.Log.InfoMsgf("版本号:%s", cfg.Version)
	} else {
		logger.Log.InfoMsgf("版本号:%s", cfg.Version)
	}
	return nil
}
