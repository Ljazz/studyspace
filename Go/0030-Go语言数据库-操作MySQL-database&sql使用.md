<!-- TOC -->

- [1.Go操作MySQL](#1go操作mysql)
  - [1.1 连接](#11-连接)
    - [1、下载依赖](#1下载依赖)
    - [2、使用MySQL驱动](#2使用mysql驱动)
    - [3、初始化连接](#3初始化连接)
    - [4、SetMaxOpenConns](#4setmaxopenconns)
    - [5、SetMaxIdleConns](#5setmaxidleconns)
  - [1.2 CRUD](#12-crud)
    - [1、建库建表](#1建库建表)
    - [2、查询](#2查询)
      - [单行查询](#单行查询)
      - [多行查询](#多行查询)
    - [插入数据](#插入数据)
    - [更新数据](#更新数据)
    - [删除数据](#删除数据)
  - [1.3 MySQL预处理](#13-mysql预处理)
    - [预处理](#预处理)
    - [预处理的好处](#预处理的好处)
    - [Go实现MySQL预处理](#go实现mysql预处理)
    - [SQL注入问题](#sql注入问题)
  - [1.4 Go实现MySQL事务](#14-go实现mysql事务)
    - [事务是什么？](#事务是什么)
    - [事务ACID](#事务acid)
    - [事务相关方法](#事务相关方法)

<!-- /TOC -->

# 1.Go操作MySQL

## 1.1 连接

Go中`database/sql`包提供了保证SQL或类SQL数据的泛用接口，并不提供具体的数据库驱动。使用`database/sql`包时必须注入（至少）一个数据库驱动。

[MySQL驱动](https://github.com/go-sql-driver/mysql)

### 1、下载依赖

> go get -u github.com/go-sql-driver/mysql

### 2、使用MySQL驱动

> `func Open(driverName, dataSourceName string) (*DB, error)`

Open打开一个dirverName指定的数据库，dataSourceName指定数据源，一般至少包含数据库文件名和其它连接必要的信息。

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // DSN: Data Source Name
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }
    defer db.Close() // 关闭要写在上面err判断的下面
}
```

### 3、初始化连接

Open函数可能只是验证其参数格式是否正确，实际上并不创建于数据库的连接。如果要检查数据源的名称是否真实有效，应该调用Ping方法。

返回的DB对象可以安全的被多个goroutine并发使用，并且维护其自己的空闲连接池。因此，Open函数应该仅被调用一次，很少需要关闭这个DB对象。

```go
// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
    // DSN:Data Source Name
    dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
    // 不会校验账号密码是否正确
    // 注意：这里不要用 := ，因为在全局已经声明
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        return err
    }
    // 尝试于数据库建立连接（校验dsn是否正确
    err = db.Ping()
    if err != nil {
        return err
    }
    return nil
}

func main() {
    err := initDB() // 调用输出化数据库的函数
    if err != nil {
        fmt.Printf("init db failed, err:%v\n", err)
        return
    }
}
```

其中`sql.DB`是表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个goroutine同时使用。

### 4、SetMaxOpenConns

> `func (db *DB) SetMaxOpenConns(n int)`

`SetMaxOpenConns`设置于数据库建立连接的最大数目。如果n大于0且小于最大闲置连接数，会将最大闲置连接数减少到匹配最大开启连接数的限制。如果n小于等于0，不会限制最大开启连接数，默认为0（无限制）。

### 5、SetMaxIdleConns

> `func (db *DB) SetMaxIdleConns(n int)`

`SetMaxIdleConns`设置连接池中的最大闲置连接数。如果n大于最大开启连接数，则新的最大闲置连接数会减少到匹配最大开启连接数的限制。如果n小于等于0，不会保留闲置连接。

## 1.2 CRUD

### 1、建库建表

```bash
mysql> CREATE DATABASE sql_test; # 创建数据库
mysql> use sql_test; # 进入数据库
Database changed
mysql> CREATE TABLE user (
    -> id BIGINT(20) NOT NULL AUTO_INCREMENT,
    -> name VARCHAR(20) DEFAULT '',
    -> age INT(11) DEFAULT '0',
    -> PRIMARY KEY(id)
    -> )engine=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
Query OK, 0 rows affected (0.03 sec)
mysql> 
```

### 2、查询

首先定义一个结构体来存储查询到的数据

```go
type user struct {
    id   int
    age  int
    name string
}
```

#### 单行查询

单行查询`db.QueryRow()`执行一次查询，并期望返回最多一行结果(Row)。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。

> `func (db *DB) QueryRow(query string, args ...interface{}) *Row`

```go
// 查询单挑数据示例
func queryRowDemo(id int) {
    sqlStr := "select id, name, age from user where id=?"
    var u user
    // 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
    err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
    if err != nil {
        fmt.Printf("scan failed, err:%v\n", err)
        return
    }
    fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}
```

#### 多行查询

多行查询`db.Query()`执行一次查询，返回多行结果(Rows)，一般用于执行select命令。参数args表示query中的占位参数。

> `func (db *DB) Query(query string, args ...interface{})(*Rows, error)`

```go
// 查询多条数据示例
func queryMultiRowDemo(n int) {
    sqlStr := "select id, name, age from user where id > ?"
    rows, err := db.Query(sqlStr, n)
    if err != nil {
        fmt.Printf("query failed, err:%v\n", err)
        return
    }
    // 关闭rows释放持有的数据库连接
    defer rows.Close()

    // 循环读取结果集中的数据
    for rows.Next() {
        var u user
        err := rows.Scan(&u.id, &u.name, &u.age)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            return
        }
        fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
    }
}
```

### 插入数据

插入、更新和删除操作都是用`Exec`方法。

> `func (db *DB) Exec(query string, args ...interface{}) (Result, err)`

Exec执行一次命令（包括查询、删除、更新、插入等），返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。

```go
func insertRowDemo() {
    sqlStr := `insert into user(name, age) values(?, ?)`
    ret, err := db.Exec(sqlStr, "王麻子", 20)
    if err != nil {
        fmt.Printf("insert failed, err:%v\n", err)
        return
    }
    id, err := ret.LastInsertId() // 新插入数据id
    if err != nil {
        fmt.Printf("get lastinsert ID failed, err:%v\n", err)
        return
    }
    fmt.Printf("insert success, the id is %d.\n", id)
}
```

### 更新数据

```go
func updateRowDemo() {
    sqlStr := "update user set age=? where id=?"
    ret, err := db.Exec(sqlStr, 39, 3)
    if err != nil {
        fmt.Printf("update failed, err:%v\n", err)
        return
    }
    n, err := ret.RowsAffected() // 操作受影响的行数
    if err != nil {
        fmt.Printf("get RowsAffected failed, err:%v\n", err)
        return
    }
    fmt.Printf("udpate success, affected rows:%d\n", n)
}
```

### 删除数据

```go
func deleteRowDemo() {
    sqlStr := "delete from user where id=?"
    ret, err := db.Exec(sqlStr, 3)
    if err != nil {
        fmt.Printf("delete failed, err:%v\n", err)
        return
    }
    n, err := ret.RowsAffected() // 操作受影响的行数
    if err != nil {
        fmt.Printf("get RowsAffected failed, err:%v\n", err)
        return
    }
    fmt.Printf("udpate success, affected rows:%d\n", n)
}
```

## 1.3 MySQL预处理

### 预处理

普通SQL语句执行过程：

1. 客户端对SQL语句进行占位符替换得到完整的SQL语句
2. 客户端发送完整的SQL语句到MySQL服务端
3. MySQL服务端执行完整的SQL语句并将结果返回给客户端

预处理执行过程：

1. 把SQL语句分成两部分，命令部分与数据部分。
2. 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理
3. 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换
4. MySQL服务端执行完整的SQL语句并将结果返回给客户端

### 预处理的好处

1. 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本
2. 避免SQL注入问题。

### Go实现MySQL预处理

`database/sql`中使用下面的`Prepare`方法来实现预处理操作。

> `func (db *DB) Prepare(query string) (*Stmt, error)`

`Prepare`方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。

查询操作的预处理示例代码如下：

```go
func prepareQueryDemo() {
    sqlstr := `select id, name, age from user where id > ?`
    stmt, err := db.Prepare(sqlStr)
    if err != nil {
        fmt.Printf("prepare failed, err:%v\n", err)
        return
    }
    defer stmt.Close()
    rows, err := stmt.Query(0)
    if err != nil {
        fmt.Printf("query failed, err:%v\n", err)
        return
    }
    defer rows.Close()
    // 循环读取数据
    for rows.Next() {
        var u user
        err := rows.Scan(&u.id, &u.name, &u.age)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            return
        }
        fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
    }
}
```

插入、更新和删除操作的预处理十分类似，代码示例如下：

```go
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
```

### SQL注入问题

任何时候都不应该自己拼接SQL

```go
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
```

此时以下输入字符串都可以引发SQL注入问题：

```go
sqlInjectDemo("xxx' or 1=1#")
sqlInjectDemo("xxx' union select * from user #")
sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
```

不同数据库中，SQL语句使用的占位符

| 数据库 | 占位符语法 |
| --- | --- |
| MySQL | ? |
| PostgreSQL | $1, $2等 |
| SQLite | ?和$1 |
| Oracle | :name |

## 1.4 Go实现MySQL事务

### 事务是什么？

**事务**：一个最下的不可再分的工作单元；通常一个事务对应一个完整的业务，同时这个完整的业务需要执行多次的DML（insert、update、delete）语句共同联合完成。

MySQL中只有使用了`Innodb`数据库引擎的数据库或表才支持事务。事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么完整的全部执行，要么全部不执行。

### 事务ACID

事务通常必须满足的4个条件：

- 原子性（Atomicity，不可分割性）
- 一致性（Consistency）
- 隔离性（Isolation，独立性）
- 持久性（Durability）

| 条件 | 解释 |
| --- | --- |
| 原子性 | 一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。 |
| 一致性 | 在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。 |
| 隔离性 | 数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。 |
| 持久性 | 事务处理结束后，对数据的修改 |

### 事务相关方法

开始事务：

> `func (db *DB) Begin() (*Tx, error)`

提交事务

> `func (tx *Tx) Commit() error`

回滚事务

> `func (tx *Tx) Rollback() error`

```go
func transactionDemo() {
    // 开启事务
    tx, err := db.Begin()
    if err != nil {
        fmt.Printf("begin failed, err:%v\n", err)
        return
    }
    // 执行多个SQL事务操作
    sqlStr1 := `update user set age=age-2 where id=1`
    sqlStr2 := `update user set age=age+2 where id=2`
    // 执行SQL1
    _, err = tx.Exec(sqlStr1)
    if err != nil {
        // 要回滚
        tx.Rollback()
        fmt.Println("执行SQL1出错了，要回滚！")
        return
    }
    // 执行SQL2
    _, err = tx.Exec(sqlStr2)
    if err != nil {
        // 要回滚
        tx.Rollback()
        fmt.Println("执行SQL2出错了，要回滚！")
        return
    }
    // 上面两步执行成功，提交本次事务
    err = tx.Commit()
    if err != nil {
        // 要回滚
        tx.Rollback()
        fmt.Println("提交出错了，要回滚！")
        return
    }
    fmt.Println("事务执行成功")
}
```
