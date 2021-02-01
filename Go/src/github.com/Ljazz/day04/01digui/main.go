package main

import "fmt"

// 递归：函数自己调用自己

func Factorial(n uint64) (result uint64) {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func fibonacci(n uint64) uint64 {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	ret := Factorial(5)
	fmt.Println(ret)
}
