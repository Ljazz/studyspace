package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	// 数据库信息
	// 用户名:密码@tcp(ip:端口)/数据库名字
	dsn := "root:walle@tcp(47.93.11.106:3306)/go"
	// 连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}

	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

// 查询数据
func queryOne(id int) {
	// 查询单条记录的sql语句
	sqlStr := `select id, name, age from user where id=?;`
	// 执行
	var u user
	// for i := 0; i < 11; i++ {
	// 	db.QueryRow(sqlStr, 1)
	// 	fmt.Printf("开始第%d次查询\n", i)
	// }

	// rowObj := db.QueryRow(sqlStr)
	// 拿到结果
	// rowObj.Scan(&u.id, &u.name, &u.age)

	db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	// 打印结果
	fmt.Printf("u:%#v\n", u)
}

// 查询多条语句
func queryMore(n int) {
	// SQL语句
	sqlStr := `select id, name, age from user where id > ?`
	// 2. 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec query failed, err:%v\n", err)
		return
	}
	// 3. 关闭rows
	defer rows.Close()
	// 4. 取值
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		fmt.Printf("u:%#v\n", u)
	}
}

// 插入
func insert() {
	// 1.写SQL语句
	sqlStr := `insert into user(name, age) values("xxx", 18);`
	// 2. exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 如果插入数据的操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed, err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

// 更新
func update() {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, 38, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

// 删除
func delete() {
	sqlStr := "delete from user where id=?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据\n", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into user(name,age) values(?,?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 后续只需要拿着stmt执行某些操作
	var m = map[string]int{
		"小龙女": 28,
		"杨过":  30,
		"李忆昔": 20,
	}
	for k, v := range m {
		_, err := stmt.Exec(k, v)
		if err != nil {
			fmt.Printf("insert data failed, err:%v\n", err)
			return
		}
	}
}

// sql注入
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功！")
	// queryOne(1)
	// queryMore(0)
	// insert()
	// update()
	// delete()
	// prepareInsert()
	sqlInjectDemo("xxx' or 1=1#")
}
