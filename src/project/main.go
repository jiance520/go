package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var myTemplate *template.Template

//type Result struct {
//	output string
//}

//结构体Result实现的方法，给自己的属性赋值
//func (p *Result) Write(b []byte) (n int, err error) {
//	fmt.Println("called by template")
//	p.output += string(b)
//	return len(b), nil
//}

type Person struct {
	Name  string
	Title string
	Age   int
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	//fmt.Fprintf(w, "hello ")
	var arr []Person
	p := Person{Name: "Mary001", Age: 10, Title: "我的个人网站"}
	p1 := Person{Name: "Mary002", Age: 10, Title: "我的个人网站"}
	p2 := Person{Name: "Mary003", Age: 10, Title: "我的个人网站"}
	arr = append(arr, p)
	arr = append(arr, p1)
	arr = append(arr, p2)

	//resultWriter := &Result{} //r:=Result{} 对结构体进行初始化，相当于new对象时，使用构造函数传参赋值，可以是结构体也可以是指针
	//io.WriteString(resultWriter, "hello world")//给result.output赋值,

	err := myTemplate.Execute(w, arr)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("template render data:", resultWriter.output)
	//myTemplate.Execute(w, p)
	//myTemplate.Execute(os.Stdout, p)
	//file, err := os.OpenFile("C:/test.log", os.O_CREATE|os.O_WRONLY, 0755)
	//if err != nil {
	//	fmt.Println("open failed err:", err)
	//	return
	//}

}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main() {
	initTemplate("D:/workspace/go/src/project/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
