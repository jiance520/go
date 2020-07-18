package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello ")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "login ") //访问页面lhttp://localhost:8880/user/login，会发送信息login给页面。
}

func history(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "history ")
}

func main() {
	//启动服务后，输入localhost:8880/跟不同的路径，会调用不同的函数(相当于控制层的cation)
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", login)
	http.HandleFunc("/user/history", history)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
