package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context
var wg sync.WaitGroup
var notify bool

func f() {
	defer wg.Done()
	for {
		fmt.Println("hello world")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	notify = true
	wg.Wait()
	// 如何通知子goroutine退出
}
