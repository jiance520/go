package main

import (
	"bufio" //缓存
	"fmt"
	"net" //tcp网络通信
	"os"  //文件
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()                       //最后关闭开启的Dial
	inputReader := bufio.NewReader(os.Stdin) //标准流读入缓存
	for {
		input, _ := inputReader.ReadString('\n') //按换行符读取，每次读一行
		trimmedinput := strings.Trim(input, "\r\n")
		if trimmedinput == "Q" { //读取到末尾，就是Q，退出quit
			return
		}
		_, err = conn.Write([]byte(trimmedinput))
		if err != nil {
			return
		}
	}
}
