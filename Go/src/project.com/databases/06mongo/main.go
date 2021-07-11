package main

import (
	"context"
	"fmt"
	"log"
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
