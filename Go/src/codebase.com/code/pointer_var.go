package main

import "fmt"

// func main(){
// 	var a int = 20 	/* 声明实际变量 */
// 	var ip *int 	/* 声明指针变量 */

// 	ip = &a 		/* 指针变量的存储地址 */

// 	fmt.Printf("a 变量的地址是：%x\n", &a)

// 	/* 指针变量的存储地址 */
// 	fmt.Printf("ip 变量的存储地址：%x\n", ip)

// 	/* 使用指针访问值 */
// 	fmt.Printf("*ip 变量的值：%d\n", *ip)
// }

// type name int8
// type first struct {
// 	a int
// 	b bool
// 	name
// }

// func main() {
// 	a := new(first)
// 	a.a = 1
// 	a.name = 11
// 	fmt.Println(a.b, a.a, a.name)
// }


// type name int8
// type first struct {
// 	a int
// 	b bool
// 	name
// }

// func main() {
// 	var a = first{1, false, 2}
// 	var b *first = &a
// 	fmt.Println(a.b, a.a, a.name, &a, b.a, &b, (*b).a)
// }

func main(){
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)
}
