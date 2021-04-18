package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	value int64
}

type Result struct {
	Job *Job
	sum int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func make_random_num(ch1 chan<- *Job) {
	defer wg.Done()
	// 循环生成int64类型的随机数，发送到jobChan
	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		ch1 <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func get_value(ch1 <-chan *Job, resultChan chan<- *Result) {
	defer wg.Done()
	for {
		job := <-ch1
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}

		newResult := &Result{
			Job: job,
			sum: sum,
		}
		resultChan <- newResult

	}
}

func main() {
	wg.Add(1)
	go make_random_num(jobChan)
	// 开启24个

	wg.Add(24)
	for i := 0; i < 24; i++ {
		go get_value(jobChan, resultChan)
	}

	for result := range resultChan {
		fmt.Printf("value:%v, sum:%v\n", result.Job.value, result.sum)
	}
	wg.Wait()
}
