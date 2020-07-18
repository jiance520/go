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
	_, err := Db.Exec("delete from person where user_id=?", 1)

	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("delete succ:")
}
