package main

import "fmt"

// array
func main() {
	// 数组初始化
	// var testArray [3]int                        // 数组回初始化为int类型的零值
	// var numArray = [3]int{1, 2}                 // 使用指定的初始值完成初始化
	// var cityArray = [3]string{"北京", "上海", "深圳"} // 使用指定的初始值完成初始化
	// fmt.Println(testArray)                      // [0 0 0]
	// fmt.Println(numArray)                       // [1 2 0]
	// fmt.Println(cityArray)                      // [北京 上海 深圳]

	// var testArray [3]int
	// var numArray = [...]int{1, 2}
	// var cityArray = [...]string{"北京", "上海", "深圳"}
	// fmt.Println(testArray)                           // [0 0 0]
	// fmt.Println(numArray)                            // [1 2]
	// fmt.Printf("type of numArray: %T\n", numArray)   // type of numArray: [2]int
	// fmt.Println(cityArray)                           // [北京 上海 深圳]
	// fmt.Printf("type of cityArray: %T\n", cityArray) // type of cityArray: [3]string

	// a := [...]int{1: 1, 3: 5}
	// fmt.Println(a)                // [0 1 0 5]
	// fmt.Printf("type of a:%T", a) // type of a:[4]int

	// 数组遍历
	// var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// 方法2：for range遍历
	// for index, value := range a {
	// 	fmt.Println(index, value)
	// }

	// 二维数组
	// a := [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	// for _, v1 := range a {
	// 	for _, v2 := range v1 {
	// 		fmt.Printf("%s\t", v2)
	// 	}
	// 	fmt.Println()
	// }

	// 练习题
	a1 := [...]int{1, 3, 5, 7, 8}
	var sum = 0
	for i := 0; i < len(a1); i++ {
		sum += a1[i]
	}
	fmt.Println("a1数组的和：", sum)
	sum = 0
	for _, value := range a1 {
		sum += value
	}
	fmt.Println("a1数组的和：", sum)

	// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	a2 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a2); i++ {
		for j := i + 1; j < len(a2); j++ {
			target := 8 - a2[i]
			if a2[j] == target {
				fmt.Printf("(%d, %d)", i, j)
			}
		}
	}
}
