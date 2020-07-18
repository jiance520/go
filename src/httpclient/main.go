package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.baidu.com") //get 8k大小
	if err != nil {
		fmt.Println("get err:", err.Error())
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data err:", err)
	}
	fmt.Println(string(data))
}
