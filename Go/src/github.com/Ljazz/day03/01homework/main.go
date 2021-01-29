package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 1. 判断字符串中汉字的数量
	s1 := "hello 天朝"
	count := 0
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)
}
