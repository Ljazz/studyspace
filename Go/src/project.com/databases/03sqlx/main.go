package main

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func initDB() (err error) {
	// 数据库信息
	// 用户名:密码@tcp(ip:端口)/数据库名字
	dsn := "root:Mr.m@2021@tcp(47.93.11.106:3306)/go"
	// 连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

type user struct {
	Id   int
	Name string
	Age  int
}

// 查询单行
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

// 查询多行
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id>?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

func insertRowDemo() {
	sqlStr := "insert into user(name, age) values(?, ?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 29)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	id, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert sucess, the id is %d.\n", id)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作受影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id=?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func insertUserDemo() {
	sqlStr := "insert into user(name, age) values(:name, :age)"
	ret, err := db.NamedExec(sqlStr, map[string]interface{}{
		"name": "七米",
		"age":  29,
	})
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	id, _ := ret.LastInsertId()
	n, _ := ret.RowsAffected()
	fmt.Printf("id:%d n:%d\n", id, n)
}

func namedQuery() {
	sqlStr := "select * from user where name=:name"
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七米"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
	u := user{
		Name: "七米",
	}
	// 使用结构体命名查询，根据结构体字段的db tag进行映射
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

func transactionDemo() {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil;don't change it
		} else {
			err = tx.Commit() // err is nil; if commit returns error update err
		}
	}()

	sqlStr1 := "update user set age=20 where id=?"
	rs, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := rs.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	if n != 1 {
		err = errors.New("exec sqlStr1 failed")
		fmt.Println(err)
		return
	}

	sqlStr2 := "update user set age=50 where i=?"
	rs, err = tx.Exec(sqlStr2, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = rs.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	if n != 1 {
		err = errors.New("exec sqlStr2 failed")
		fmt.Println(err)
		return
	}

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	queryRowDemo()
	fmt.Println("****************")
	queryMultiRowDemo()
	fmt.Println("****************")
	insertRowDemo()
	fmt.Println("****************")
	// updateRowDemo()
	fmt.Println("****************")
	// deleteRowDemo()
	fmt.Println("****************")
	insertUserDemo()
	fmt.Println("****************")
	namedQuery()
}
