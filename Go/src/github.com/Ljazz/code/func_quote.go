package main

import "fmt"

// 实现参数 +1 操作
func add1(a *int) int { // 参数使用指针类型
	*a = *a + 1 // 修改a的值
	return *a   // 返回新的值
}

func main() {
	x := 3
	fmt.Println("x = ", x)
	x1 := add1(&x)
	fmt.Println("x1 = ", x1)
	fmt.Println("x = ", x)
}
