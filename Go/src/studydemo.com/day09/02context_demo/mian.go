package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context
var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("hello world")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	exitChan <- true
	wg.Wait()
	// 如何通知子goroutine退出
}
