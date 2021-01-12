package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 &a = %x a = %d\n", &a, a)
	fmt.Printf("指针变量 ptr = %x *ptr = %d &ptr = %x\n", ptr, *ptr, &ptr)
	fmt.Printf("指向指针的指针变量 pptr = %x **pptr = %d\n", pptr, **pptr)
}
