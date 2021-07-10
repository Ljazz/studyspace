package main

import "fmt"

// 自定义类型和类型别名

// type后面跟的是类型
type MyInt int     // 自定义类型
type yourInt = int // 类型别名

func main() {
	var n MyInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)
}
