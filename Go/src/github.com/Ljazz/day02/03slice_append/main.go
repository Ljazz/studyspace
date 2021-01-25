package main

import "fmt"

// append() 为切片添加元素

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	// 调用append函数必需用原来的切片变量接收返回值
	// append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	// 必需用变量接收append的返回值
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) // ... 表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	v := a1[:]

	// 删除搜索引为1的
	v = append(v[0:1], v[2:]...)
	fmt.Println(v)
	fmt.Println(a1)
}
