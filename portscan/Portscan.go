package poarscan

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

//var sem = make(chan struct{}, 100)

func scan(ip string, startport) {
	defer wg.Done()
	//sem <- struct{}{}
	//defer func() { <-sem }()
	for i := startport; i <= endport; i++ {

	}
	Address := fmt.Sprintf("%s:%d", ip, port)                   //ip端口
	conn, err := net.DialTimeout("tcp", Address, 2*time.Second) //2秒
	if err != nil {
		return
	} else {
		fmt.Printf("端口%d开放\n", port)
	}
	defer conn.Close()
}

func main() {

	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go scan("127.0.0.1", i)
	}
	wg.Wait()
	fmt.Println("扫描结束")
}
