package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //_表示初始化mysql驱动，但不会在代码中调用此包的内容
	"github.com/jmoiron/sqlx"          //增删改查
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}
func main() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")

	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	//查看id，来检查是否插入成功
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}
