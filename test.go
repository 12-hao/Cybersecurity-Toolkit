package main

import (
	"fmt"
	"github.com/imroc/req/v3"
)

func main() {

	c := req.C()
	r := c.R()
	get, err := r.Get("www.bing.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(get.TotalTime())
}
