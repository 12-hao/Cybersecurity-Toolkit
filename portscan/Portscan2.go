package poarscan

import (
	"fmt"
	"net"
	"time"
)

func pscan(ip string, port int, method string, ch chan bool) {
	address := fmt.Sprintf("%s:%d", ip, port)                    //ip端口
	conn, err := net.DialTimeout(method, address, 2*time.Second) //2秒
	if err != nil {
		ch <- false
		return
	} else {
		ch <- true
	}
	defer conn.Close()
}

/*func main() {

	res := make(chan bool)
	for i := 0; i < 65535; i++ {
		go pscan("43.139.118.102", i, "tcp", res)
		data := <-res
		//这里用通道的作用是因为这里是无缓冲通道 这个通道会导致阻塞直到接受方把里面的数据拿走，用来使得主线程等待协程的作用
		if data == true {
			fmt.Printf("%d端口开放", i)
		}
	}
	fmt.Println("扫描结束")
}*/
