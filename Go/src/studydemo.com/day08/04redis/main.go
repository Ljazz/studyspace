package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// redis

var redisDB *redis.Client

func initRedis() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "47.93.11.106:6379",
		Password: "blueprint@2021", // password
		DB:       0,                // use default DB
		PoolSize: 100,              // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = redisDB.Ping(ctx).Result()
	return err
}

func RedisOperation() {
	ctx := context.Background()
	if err := initRedis(); err != nil {
		fmt.Printf("Redis连接失败，错误信息：%v", err)
		return
	}
	fmt.Println("Redis连接成功")

	err := redisDB.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisDB.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := redisDB.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
