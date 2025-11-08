package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup
var sem = make(chan struct{}, 100)

type Result struct {
	Idr    string `json:"ip_address"`    //IP地址
	Portrr int    `json:"port"`          //对应端口
	Vulr   string `json:"vulnerability"` //端口对应功能名称
	Serr   string `json:"severity"`      //安全等级
	Timer  string `json:"timestamp"`     //扫描的时间
}

var flag int = 0
var results []Result

func Scan(ip string, port int) {
	defer wg.Done()
	sem <- struct{}{}
	defer func() { <-sem }()
	Address := fmt.Sprintf("%s:%d", ip, port) //ip端口
	conn, err := net.DialTimeout("tcp", Address, 3*time.Second)
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
	db, err := sql.Open("sqlite3", "./Nmap.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS scan_results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT NOT NULL,
		port INT NOT NULL,
		vulnerability TEXT NOT NULL,
		severity TEXT NOT NULL,
		timestamp TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println("创建表失败:", err)
		return
	}

	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go Scan("127.0.0.1", i)
	}
	wg.Wait()
	fmt.Println("扫描结束")
	for i := 0; i < len(results); i++ {

	}
	for i := 0; i < len(results); i++ {
		intosql := `INSERT Into scan_results (ip, port, vulnerability, severity, timestamp) values (?,?,?,?,?)`
		_, err := db.Exec(intosql, results[i].Idr, results[i].Portrr, results[i].Vulr, results[i].Serr, results[i].Timer)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("写入文件成功")
}
