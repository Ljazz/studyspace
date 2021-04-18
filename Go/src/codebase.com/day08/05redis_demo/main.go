package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "47.93.11.106:6379", redis.DialPassword("blueprint@2021"))
	if err != nil {
		fmt.Println("conn redis failed, ", err)
		return
	}
	defer c.Close()
	fmt.Println("redis conn success")

	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println("Set failed, ", err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed, ", err)
		return
	}
	fmt.Println(r)
}
