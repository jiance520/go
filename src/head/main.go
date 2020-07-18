package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {
	for _, v := range url {
		c := http.Client{ //鼠标悬停查看，http包里的结构体Client{}，引入的包相当于一个实例
			Transport: &http.Transport{ //字面量给结构体Client的接口Transport赋值。赋值同类型&http.Transport
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2 //延时不超过2秒，否则失败
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}

		fmt.Printf("head succ, status:%v\n", resp.StatusCode) //状态200成功，resp响应包括所有内容
	}
}
