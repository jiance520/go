package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //_表示初始化mysql驱动，但不会在代码中调用此包的内容
	"github.com/jmoiron/sqlx"          //增删改查
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

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
	var person []Person
	err := Db.Select(&person, "select * from person where user_id=?", 1)

	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)
}
