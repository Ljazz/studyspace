package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context
var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("hello world")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	// context.Background() 根
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	cancel()
	wg.Wait()
	// 如何通知子goroutine退出
}
