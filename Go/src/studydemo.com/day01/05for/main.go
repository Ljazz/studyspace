package main

import "fmt"

// for循环

func main() {
	// 基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种1
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// 变种2
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for range循环
	// s := "hello world"
	// for i, v := range s {
	// 	fmt.Printf("%d %c\n", i, v)
	// }

	// 9x9乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d  ", j, i, i*j)
		}
		fmt.Println()
	}
}
