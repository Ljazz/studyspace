# Go语言操作mongoDB

## mongoDB介绍

[mongoDB](https://www.mongodb.com/)是目前比较流行的一个基于分布式文件存储的数据库，它是一个介于关系数据库和非关系数据库(NoSQL)之间的产品，是非关系数据库中功能最丰富，最像关系数据库的。

mongoDB中将一条数据存储为一个文档（document），数据结构由键值(key-value)对组成。其中文档类似于我们平常编程中用到的JSON对象。文档中的字段值可以包含其他文档，数组及文档数组。

### mongoDB相关概念

| MongoDB术语/概念 | 说明 | 对比SQL术语/概念 |
| --- | --- | --- |
| database | 数据库 | database |
| collection | 集合 | table |
| document | 文档 | row |
| field | 字段 | column |
| index | index | 索引 |
| primary key | 主键MongoDB自动将_id字段设置为主键 | primary key |

## mongoDB安装

官网下载地址：https://www.mongodb.com/try/download/community

官方安装教程：https://docs.mongodb.com/manual/administration/install-community/

## mongoDB基本使用

### 启动mongoDB数据库

$\color{red}{Windows}$

> "C:\Program Files\MongoDB\Server\4.2\bin\mongod.exe" --dbpath="c:\data\db"

$\color{red}{Mac}$

> mongod --config /usr/local/etc/mongod.conf

or
> brew services start mongodb-community@4.2

### 启动client

$\color{red}{Windows}$

> "C:\Program Files\MongoDB\Server\4.2\bin\mongod.exe"

$\color{red}{Mac}$

> mongod

### 数据库常用命令

`show dbs;`：查看数据库

`use q1mi;`：切换到指定数据库，如果不存在该数据库就创建。

`db;`：显式当前所在数据库

`db.dropDatabase();`：删除当前数据库

### 数据集常用命令
`db.createCollection(name, options)`：创建数据集
- name：数据集名称
- options：可选参数，指定内存大小和索引

`show collections;`：查看当前数据库中所有集合

`db.student.drop();`：删除指定数据集

### 文档常用命令

插入一条文档：
```bash
> db.student.insertOne({name:"小王子",age:18});
{
    "acknowledged" : true,
    "insertedId" : ObjectId("5db149e904b33457f8c02509")
}
```

插入多条文档：
```bash
> db.student.insertMany([
    {name: "张三", age:20},
    {name: "李四", age: 25}
]);
{
    "acknowledged" : true,
    "insertedIds" : [
        ObjectId("5db14c4704b33457f8c0250a"),
        ObjectId("5db14c4704b33457f8c0250b")
    ]
}
```

查看所有文档
```bash
> db.student.find();
{ "_id" : ObjectId("5db149e904b33457f8c02509"), "name" : "小王子", "age" : 18 }
{ "_id" : ObjectId("5db14c4704b33457f8c0250a"), "name" : "张三", "age" : 20 }
{ "_id" : ObjectId("5db14c4704b33457f8c0250b"), "name" : "李四", "age" : 25 }
```
查看`age>20`的文档
```bash
> db.student.find(
    {age: {$gt: 20}}
);
{ "_id" : ObjectId("5db14c4704b33457f8c0250b"), "name" : "李四", "age" : 25 }
```
更新文档
```bash
> db.student.update(
    {name: "小王子"},
    {name: "老王子", age: 98}
);
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.student.find()
{ "_id" : ObjectId("5db149e904b33457f8c02509"), "name" : "老王子", "age" : 98 }
{ "_id" : ObjectId("5db14c4704b33457f8c0250a"), "name" : "张三", "age" : 20 }
{ "_id" : ObjectId("5db14c4704b33457f8c0250b"), "name" : "李四", "age" : 25 }
```
删除文档
```bash
> db.student.deleteOne({name: "李四"});
{ "acknowledged" : true, "deletedCount" : 1 }
> db.student.find()
{ "_id" : ObjectId("5db149e904b33457f8c02509"), "name" : "老王子", "age" : 98 }
{ "_id" : ObjectId("5db14c4704b33457f8c0250a"), "name" : "张三", "age" : 20 }
```

## Go语言操作mongoDB

### 安装mongoDB Go驱动包

> go get go.mongodb.org/mongo-driver

通过Go代码链接mongoDB

```go
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 设置客户端链接配置
	clientOptions := options.Client().ApplyURI("mongodb://47.93.11.106:27017")

	// 链接到mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 链接检查
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}
```