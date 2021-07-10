package main

import "fmt"

// 变量声明
// var name string
// var age int
// var isOk bool

// 批量声明

var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "理想"
	age = 23
	isOk = true
	// var hehe string
	// Go语言中变量声明之后必须使用，不适用就编译不过去
	fmt.Print(isOk) // 在终端中输出要打印的内容
	fmt.Println()
	fmt.Printf("name: %s\n", name) // %s：占位符，使用name这个变量的值去替换占位符
	fmt.Println(age)               // 打印完指定的内容之后会在后面加一个换行符
	// hehe = "呵呵"

	// 声明变量同时赋值
	var s1 string = "whb"
	fmt.Println(s1)
	// 类型推导（根据值类型判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明
	s3 := "哈哈"
	fmt.Println(s3)
}
