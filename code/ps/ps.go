package ps

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func ScanPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return // 端口关闭或不可达
	}
	conn.Close()
	fmt.Printf("[+] Port %d is open\n", port)
}
