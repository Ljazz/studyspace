package main

import "fmt"

// map

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // 还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10) // 要估算号map容量避免程序中动态扩容
	m1["理想"] = 18
	m1["jiwuming"] = 35

	fmt.Println(m1)
	fmt.Println(m1["理想"])
	// 约定ok接收返回的布尔值
	value, ok := m1["娜扎"] // 如果不存在这个key拿到对应值类型的零值
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 遍历key
	for key := range m1 {
		fmt.Println(key)
	}
	// 遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}

	//
}
