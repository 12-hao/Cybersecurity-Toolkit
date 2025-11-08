package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var results []Result

type Result struct {
	Idr    string `json:"ip_address"`    //IP地址
	Portrr int    `json:"port"`          //对应端口
	Vulr   string `json:"vulnerability"` //端口对应功能名称
	Serr   string `json:"severity"`      //安全等级
	Timer  string `json:"timestamp"`     //扫描的时间
}
type PortScan struct {
	Ipadd      []string `json:"ip_addresses"`
	Port_range string   `json:"port_range"`
	Timeout    int      `json:"timeout"`
}

var flag int = 0

func Scan1(ip string, port int, timeout int) {
	defer wg1.Done()
	Address := fmt.Sprintf("%s:%d", ip, port) //ip端口
	conn, err := net.DialTimeout("tcp", Address, time.Duration(timeout)*time.Second)
	if err != nil {
		return
	} else {
		fmt.Printf("%s的%d端口开放\n", ip, port)
		var rank string
		var agreement string
		switch port {
		case 21: // FTP
			rank = "High"
			agreement = "FTP"
		case 22: // SSH
			rank = "High"
			agreement = "SSH"
		case 23: // Telnet
			rank = "High"
			agreement = "Telnet"
		case 25: // SMTP
			rank = "Medium"
			agreement = "SMTP"
		case 53: // DNS
			rank = "Medium"
			agreement = "DNS"
		case 80: // HTTP
			rank = "Medium"
			agreement = "Web"
		case 110: // POP3
			rank = "Medium"
			agreement = "POP3"
		case 135: // RPC
			rank = "High"
			agreement = "RPC"
		case 139: // NetBIOS
			rank = "High"
			agreement = "NetBIOS"
		case 443: // HTTPS
			rank = "Medium"
			agreement = "Web"
		case 445: // SMB
			rank = "High"
			agreement = "SMB"
		case 1433: // MSSQL
			rank = "High"
			agreement = "MSSQL"
		case 3306: // MySQL
			rank = "High"
			agreement = "MySQL"
		case 3389: // RDP (Remote Desktop Protocol)
			rank = "High"
			agreement = "RDP"
		case 5432: // PostgreSQL
			rank = "High"
			agreement = "PostgreSQL"
		case 5900: // VNC
			rank = "High"
			agreement = "VNC"
		case 8080: // HTTP Proxy
			rank = "Medium"
			agreement = "Web"
		default:
			rank = "Low"
			agreement = "Unknown"
		}
		results = append(results, Result{
			Idr:    ip,
			Portrr: port,
			Vulr:   agreement,
			Serr:   rank,
			Timer:  time.Now().Format("2006年01月02日15.04分"),
		})
	}
	defer conn.Close()
} //端口扫描函数

func main() {
	file, err := os.OpenFile("scan_config.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("1处报错", err)
		return
	}

	defer file.Close()
	Ps := PortScan{}
	all, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("2处报错", err)
		return
	}
	err = json.Unmarshal(all, &Ps)
	if err != nil {
		fmt.Println("3处报错", err)
		return
	}

	//上面这段是把json数据读入进来
	split := strings.Split(Ps.Port_range, "-") //将端口分割出来
	Portl, _ := strconv.Atoi(split[0])         //分别获取左右区间
	Portr, _ := strconv.Atoi(split[1])

	for i := 0; i < len(Ps.Ipadd); i++ { //外层循环IP
		for j := Portl; j <= Portr; j++ { //内层循环端口
			wg1.Add(1)
			go Scan1(Ps.Ipadd[i], j, Ps.Timeout) //并发执行
		}
		wg1.Wait()
		fmt.Printf("%s扫描结束\n", Ps.Ipadd[i])
	}

	fmt.Println("扫描结束")
	marshal, err := json.MarshalIndent(results, " ", "    ")
	if err != nil {
		fmt.Println("4处报错", err)
		return
	}
	openFile, err := os.OpenFile("scan_result", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("5处错误", err)
		return
	}
	defer openFile.Close()
	_, err = openFile.Write(marshal)
	if err != nil {
		fmt.Println("6处错误", err)
		return
	}
	fmt.Println("结果已输出到文件中")
}
