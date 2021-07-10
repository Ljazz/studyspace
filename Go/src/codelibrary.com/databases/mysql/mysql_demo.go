package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UesrId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var db *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:walle@tcp(47.93.11.106:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed: ", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping failed: ", err)
		return
	}

	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	defer db.Close()
}

func insert() {
	r, err := db.Exec(`insert into person(username, sex, email)values("stu001", "man", "stu01@qq.com")`)
	if err != nil {
		fmt.Println("exec failed: ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed: ", err)
		return
	}
	fmt.Println("insert succ: ", id)
}

func main() {
	insert()
}
