package main

import "fmt"

// 类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if ok {
		fmt.Println("传进来的是一个字符串：", str)
	} else {
		fmt.Println("猜错了")
	}
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch v := a.(type) {
	case string:
		fmt.Println("是一个字符串：", v)
	case int:
		fmt.Println("是一个int：", v)
	case bool:
		fmt.Println("是一个bool：", v)

	}
}

func main() {
	assign(100)
	assign2(100)
}
