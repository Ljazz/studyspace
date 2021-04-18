/***************************************************
	 001-helloworld.go
****************************************************/

// package main
// import "fmt"
// func main() {
// 	/* hello world程序 */
// 	fmt.Println("hello world!")
// }

/***************************************************
	 002-string_print.go
****************************************************/
// package main
// import "fmt"
// func main() {
// 	var temp string
// 	temp = `
// 		x := 10
// 		y := 20
// 		z := 30
// 		fmt.Println(x, " ", y, " ", z)
// 		x, y, z = y, z, x
// 		fmt.Println(x, " ", y, " ", z)`
// 	fmt.Println(temp)
// }

/***************************************************
	 004-通用打印格式.go
****************************************************/
// package main
// import "fmt"
// func main() {
// 	str := "steven"
// 	fmt.Println("%v, %+v, %#v, %T", str, str, str, str)
// 	var a rune = '一'
// 	fmt.Println("%v, %+v, %#v, %T", a, a, a, a)
// 	var b byte = 'b'
// 	fmt.Println("%v, %+v, %#v, %T", b, b, b, b)
// 	var c int32 = 98
// 	fmt.Println("%v, %+v, %#v, %T", c, c, c, c)
// }

/***************************************************
	 005-bool_print.go
****************************************************/
// package main
// import "fmt"
// func main() {
// 	var flag bool
// 	fmt.Println("%T, %t", flag, flag)
// 	flag = true
// 	fmt.Println("%T, %t", flag, flag)
// }

/***************************************************
	 006-整型打印格式
****************************************************/
//package main
//import "fmt"
//func main() {
//	fmt.Println("%T , %d ", 123, 123)
//	fmt.Println("%T , %5d ", 123, 123)
//	fmt.Println("%T , %05d ", 123, 123)
//	fmt.Println("%T , %b ", 123, 123)
//	fmt.Println("%T , %o ", 123, 123)
//	fmt.Println("%T , %c ", 123, 123)
//	fmt.Println("%T , %q ", 123, 123)
//	fmt.Println("%T , %x ", 123, 123)
//	fmt.Println("%T , %X ", 123, 123)
//	fmt.Println("%T , %U ", 123, 123)
//}

/***************************************************
	 007-浮点型打印格式
****************************************************/
// package main
// import "fmt"
// func main()  {
// 	fmt.Println("%b", 123.123456)
// 	fmt.Println("%f", 123.123456)
// 	fmt.Println("%.2f", 123.123456)
// 	fmt.Println("%e", 123.123456)
// 	fmt.Println("%E", 123.123456)
// 	fmt.Println("%.1e", 123.123456)
// 	fmt.Println("%F", 123.123456)
// 	fmt.Println("%g", 123.123456)
// 	fmt.Println("%G", 123.123456)
// }

/***************************************************
	 008-复数打印格式
****************************************************/
// package main
// import "fmt"
// func main() {
// 	var value complex64 = 2.1 + 12i
// 	value2 := complex(2.1, 12)
// 	fmt.Println(real(value))
// 	fmt.Println(imag(value))
// 	fmt.Println(value2)
// }

/***************************************************
	 009-字符串和字节数组打印格式
****************************************************/
// package main
// import "fmt"
// func main() {
// 	arr := [] byte{'x', 'y', 'z', 'z'}
// 	fmt.Println("%s", "欢迎大家学习区块链")
// 	fmt.Println("%q", "欢迎大家学习区块链")
// 	fmt.Println("%x", "欢迎大家学习区块链")
// 	fmt.Println("%X", "欢迎大家学习区块链")
// 	fmt.Println("%T, %s", arr, arr)
// 	fmt.Println("%T, %q", arr, arr)
// 	fmt.Println("%T, %x", arr, arr)
// 	fmt.Println("%T, %X", arr, arr)
// }

/***************************************************
	 010-类型转换
****************************************************/
// package main

// import "fmt"

// func main() {
// 	var a int = 100
// 	b := float64(a)
// 	c := string(a)
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

/***************************************************
	 010-浮点型和整型之间转换
****************************************************/
//package main
//
//import "fmt"
//
//func main() {
//	chinese := 90
//	english := 80.9
//	//avg := (chinese + english) / 2
//	avg2 := (float64(chinese) + english) / 2
//	//fmt.Println("%T, %d", avg, avg)
//	fmt.Println("%T, %f", avg2, avg2)
//}

// package main

// import "fmt"

// func main() {
// 	/* 使用if...else语句判断奇数偶数 */
// 	num := 20
// 	if num%2 == 0 {
// 		fmt.Println(num, "偶数")
// 	} else {
// 		fmt.Println(num, "奇数")
// 	}
// 	/* 上述代码运行结果为：num 偶数 */

// 	/* if...else if ...else判断学生成绩 */
// 	score := 88
// 	if score >= 90 {
// 		fmt.Println("优秀")
// 	} else if score >= 80 {
// 		fmt.Println("良好")
// 	} else if score >= 70 {
// 		fmt.Println("中等")
// 	} else if score >= 60 {
// 		fmt.Println("及格")
// 	} else {
// 		fmt.Println("不及格")
// 	}
// 	/* 上述代码运行结果为：良好 */
// }

// package main
// import "fmt"
// func main()  {
// 	if score := 98; score >= 60 {
// 		if score >= 70 {
// 			if score >= 80 {
// 				if score >= 90 {
// 					fmt.Println("优秀")
// 				}else{
// 					fmt.Println("良好")
// 				}
// 			}else{
// 				fmt.Println("中等")
// 			}
// 		} else {
// 			fmt.Println("及格")
// 		}
// 	} else{
// 		fmt.Println("不及格")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	/* 定义全局变量 */
// 	grade := ""
// 	score := 78.5
// 	switch { // switch后面省略不写，默认相当于：switch true
// 	case score >= 90:
// 		grade = "A"
// 	case score >= 80:
// 		grade = "B"
// 	case score >= 70:
// 		grade = "C"
// 	case score >= 60:
// 		grade = "D"
// 	default:
// 		grade = "E"
// 	}
// 	fmt.Println("你的等级是：", grade)
// 	fmt.Print("最终评价是：")
// 	switch grade {
// 	case "A":
// 		fmt.Println("优秀")
// 	case "B":
// 		fmt.Println("良好")
// 	case "C":
// 		fmt.Println("中等")
// 	case "D":
// 		fmt.Println("及格")
// 	default:
// 		fmt.Println("差")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var x interface{}
// 	switch i := x.(type) {
// 	case nil:
// 		fmt.Println("x的类型：%T", i)
// 	case int:
// 		fmt.Println("x 是 int 型")
// 	case float64:
// 		fmt.Println("x 是 float64 型")
// 	case func(int) float64:
// 		fmt.Println("x 是 func(int) 型")
// 	case bool, string:
// 		fmt.Println("x 是 bool或string 型")
// 	default:
// 		fmt.Println("位置型")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var i int
// 	for i <= 10 {
// 		fmt.Print(i)
// 		i++
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	/* 1~100的和 */
// 	sum := 0
// 	for i := 1; i <= 100; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)

// 	/* 1~40所有3的倍数 */
// 	j := 1
// 	num := 0
// 	for j <= 40 {
// 		if j%3 == 0 {
// 			num += j
// 			fmt.Print(j)
// 			if j < 39 {
// 				fmt.Print("+")
// 			} else {
// 				fmt.Print(" = ", num)
// 			}
// 		}
// 		j++
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	//定义行数
// 	lines := 8
// 	for i := 0; i < lines; i++ {
// 		for n := 0; n < 2*i+1; n++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

//package main
//
//import "fmt"
//
//func main() {
//	var C, c int //声明变量
//	C = 1
//LOOP:
//	for C < 50 {
//		C++ //C=1不能写入for
//		for c = 2; c < C; c++ {
//			if C%c == 0 {
//				goto LOOP //若发现因子，则不是素数
//			}
//		}
//		fmt.Println(C)
//	}
//}
//package main

///* 生命全局变量 */
//var a1 int = 7
//var b1 int = 9
//func main() {
//	/* main函数中声明局部变量 */
//	a1, b1, c1 := 10, 20, 0
//	fmt.Printf("main()函数中 a1 = %d\n", a1)
//	fmt.Printf("main()函数中 b1 = %d\n", b1)
//	fmt.Printf("main()函数中 c1 = %d\n", c1)
//	c1 = sum(a1, b1)
//	fmt.Printf("main()函数中 c1 = %d\n", c1)
//}
//
///* 函数定义-两数相加 */
//func sum (a1, b1 int) (c1 int){
//	a1 ++
//	b1 += 2
//	c1 = a1 + b1
//	fmt.Printf("sum()函数中 a1 = %d\n", a1)
//	fmt.Printf("sum()函数中 b1 = %d\n", b1)
//	fmt.Printf("sum()函数中 c1 = %d\n", c1)
//	return c1
//}

//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	result := StringToLower("AbcdefGHijklMNOPqrstUVWxyz", processCase)
//	fmt.Println(result)
//	result = StringToLower2("AbcdefGHijklMNOPqrstUVWxyz", processCase)
//	fmt.Println(result)
//}
//
//// 处理字符串，奇数偶数依次显示为大小写
//func processCase(str string) string {
//	result := ""
//	for i, value := range str{
//		if i%2 == 0{
//			result += strings.ToUpper(string(value))
//		}else {
//			result += strings.ToLower(string(value))
//		}
//	}
//	return result
//}
//func StringToLower(str string, f func(string) string) string {
//	fmt.Printf("%T \n", f)
//	return f(str)
//}
//
//type caseFunc func(string) string // 声明一个函数类型，通过type关键字，caseFunc会形成一种新的类型
//
//func StringToLower2(str string, f caseFunc) string {
//	fmt.Printf("%T \n", f) //打印变量f的类型
//	return f(str)
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main()  {
//	/* 1. 在定义时调用匿名函数 */
//	func(data int){
//		fmt.Println("hell", data)
//	}(100)
//
//	/* 2. 将匿名函数赋值给变量 */
//	f := func (data string) {
//		fmt.Println(data)
//	}
//	f("欢迎学习Go语言")
//
//	/* 3.匿名函数用作回调函数 */
//	// 调用函数，对每个元素进行求平方根操作
//	arr := [] float64 {1, 9, 16, 25, 30}
//	visit (arr, func(v float64){
//		v = math.Sqrt(v)
//		fmt.Printf("%.2f \n", v)
//	})
//	//调用函数，对每个元素进行求平方根操作
//	visit(arr, func(v float64){
//		v = math.Pow(v, 2)
//		fmt.Printf("%.0f \n", v)
//	})
//}
//// 定义一个函数，遍历切片元素，对每个切片元素进行处理
//func visit(list []float64, f func(float64)) {
//	for _, value := range list{
//		f(value)
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	// 1. 没有使用闭包进行计数
//	for i := 0; i < 5; i++{
//		fmt.Printf("i = %d\t", i)
//		fmt.Println(add(i))
//	}
//
//	// 2.使用闭包实现计数
//	pos := adder()
//	for i := 0; i< 10; i++ {
//		fmt.Printf("i = %d\t", i)
//		fmt.Println(pos(i))
//	}
//	fmt.Println("-------------------------")
//	for i := 0; i < 10; i++ {
//		fmt.Printf("i = %d\t", i)
//		fmt.Println(pos(i))
//	}
//}
//
//func add (x int) int {
//	sum := 0
//	sum += x
//	return sum
//}
//
//func adder () func(int) int {
//	sum := 0
//	return func(x int) int {
//		fmt.Printf("sum1 = %d \t", sum)
//		sum += x
//		fmt.Printf("sum2 = %d \t", sum)
//		return sum
//	}
//}

// package main

// import "fmt"

// func mian() {
// 	myfunc := Counter()
// 	fmt.Printf("%T\n", myfunc)
// 	fmt.Println("myfunc", myfunc)
// 	/* 调用myfunc函数，i变量自增1并返回 */
// 	fmt.Println(myfunc())
// 	fmt.Println(myfunc())
// 	fmt.Println(myfunc())
// 	/* 创建新的函数 next Number1，并查看结果 */
// 	myfunc1 := Counter()
// 	fmt.Println("myfunc1", myfunc1)
// 	fmt.Println(myfunc1())
// 	fmt.Println(myfunc1())
// }

// func Counter () func() int {
// 	i := 0
// 	res := func() int {
// 		i += 1
// 		return i
// 	}
// 	fmt.Printf("%T, %v \n", res, res)  // func() int ,0x1095af0
// 	fmt.Println("Counter中的内部函数：", res)  // func() int ,0x1095af0
// 	return res
// }

//package main
//
//import "fmt"
//
//func main() {
//	myfunc := Counter()
//	fmt.Printf("%T\n", myfunc) // func() int
//	fmt.Println("myfunc", myfunc) // myfunc 0x49c930
//	/* 调用myfunc函数，i变量自增1并返回 */
//	fmt.Println(myfunc()) // 1
//	fmt.Println(myfunc()) // 2
//	fmt.Println(myfunc()) // 3
//	/* 创建新的函数 next Number1，并查看结果 */
//	myfunc1 := Counter()
//	fmt.Println("myfunc1", myfunc1) // myfunc1 0x49c930
//	fmt.Println(myfunc1()) // 1
//	fmt.Println(myfunc1()) // 2
//}
//
//func Counter() func() int {
//	i := 0
//	res := func() int {
//		i += 1
//		return i
//	}
//	fmt.Printf("%T, %v \n", res, res)  // func() int, 0x49c930
//	fmt.Println("Counter中的内部函数：", res) // Counter中的内部函数： 0x49c930
//	return res
//}

/*************************************************************************
	不定长参数函数
*************************************************************************/
//package main
//
//import "fmt"
//
//func main(){
//	sum, avg, count := GetScore(90, 82.5, 73, 64.8)
//	fmt.Printf("学员共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f", count, sum, avg) // 学员共有4门成绩，总成绩为：310.30，平均成绩为：77.58
//	fmt.Println()
//	scores := []float64{92, 72.5, 93, 74.5, 89, 87, 74}
//	sum, avg, count = GetScore(scores ...)
//	fmt.Printf("学员共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f", count, sum, avg) // 学员共有7门成绩，总成绩为：582.00，平均成绩为：83.14
//}
//func GetScore(scores...float64) (sum, avg float64, count int)  {
//	for _, value := range scores {
//		sum += value
//		count++
//	}
//	avg = sum / float64(count)
//	return
//}

/*************************************************************************
	递归函数
*************************************************************************/
//package main
//
//import "fmt"
//
//func main()  {
//	fmt.Println(factorial(5)) // 120
//	fmt.Println(getMultiple(5)) // 120
//}
////使用递归实现阶乘
//func factorial(n int) int  {
//	if n == 0 {
//		return 1
//	}
//	return n * factorial(n-1)
//}
////使用循环实现阶乘
//func getMultiple(num int) (result int)  {
//	sum := 1
//	for i := 1; i <= num; i++{
//		sum *= i
//	}
//	return sum
//}

/*************************************************************************
指针
	1. 取地址
	2. 指针的使用
*************************************************************************/
//package main
//
//import "fmt"
//
//func main()  {
//	a := 10
//	fmt.Printf("变量的地址：%x", &a) // 变量的地址：c00000a0c8
//}

//package main
//
//import "fmt"
//
//func main(){
//	//声明实际变量
//	var a int = 120
//	// 声明指针变量
//	var ip *int
//	//给指针变量赋值，将变量a的地址赋值给ip
//	ip = &a
//	//打印 a 的类型和值
//	fmt.Printf("a 的类型是%T， 值是%v \n", a, a) // a 的类型是int， 值是120
//	//打印 &a 的类型和值
//	fmt.Printf("&a 的类型是%T，值是%v \n", &a, &a) // &a 的类型是*int，值是0xc000070090
//	//打印 ip 的类型和值
//	fmt.Printf("ip 的类型是%T，值是%v \n", ip, ip) // ip 的类型是*int，值是0xc000070090
//	//打印变量 *ip 的类型和值
//	fmt.Printf("*ip 的类型是%T，值是%v \n", *ip, *ip) // *ip 的类型是int，值是120
//	//打印变量 *&a 的类型和值
//	fmt.Printf("*&a 的类型是%T，值是%v \n", *&a, *&a) // *&a 的类型是int，值是120
//	fmt.Println(a, &a, *&a) // 120 0xc000070090 120
//	fmt.Println("=======================")
//	fmt.Println(ip, &ip, *ip, *(&ip), &(*ip)) // 0xc000070090 0xc00009a018 120 0xc000070090 0xc000070090
//}

//package main
//
//import "fmt"
//
//func main() {
//	// 改变变量的值
//	b := 3158
//	a := &b
//	fmt.Println("b 的地址：", a) // b 的地址： 0xc00000a0c8
//	fmt.Println("*a 的值：", *a) // *a 的值： 3158
//	*a++
//	fmt.Println("b 的地址：", b) // b 的地址： 3159
//}

//package main
//
//import "fmt"
//
//func main()  {
//	// 使用指针作为函数的参数
//	a := 58
//	fmt.Println("函数调用前 a 的值：", a) // 函数调用前 a 的值： 58
//	fmt.Printf("%T \n", a) // int
//	fmt.Printf("%x \n", &a) // c00006c090
//	// b := &a
//	var b *int = &a
//	change(b)
//	fmt.Println("函数调用后 a 的值：", a) // 函数调用后 a 的值： 17
//}
//func change(val *int)  {
//	*val = 17
//}

//package main
//
//import "fmt"
//
//const COUNT int  = 4
//func main()  {
//	// 指针数组
//	a := [COUNT]string {"abc", "ABC", "123", "一二三"}
//	i := 0
//	// 定义指针数组
//	var ptr [COUNT]*string
//	fmt.Printf("%T, %v \n", ptr, ptr)
//	for i=0; i< COUNT;i++{
//		// 将数组中每个元素的地址赋值给指针数组
//		ptr[i] = &a[i]
//	}
//	fmt.Printf("%T, %v \n", ptr, ptr)
//	//获取指针数组中第一个值，其实就是个地址
//	fmt.Println(ptr[0])
//	//根据数组元素的每个地址获取该地址所对应的元素的数值
//	for i=0;i<COUNT;i++{
//		fmt.Printf("a[%d] = %s \n", i, *ptr[i])
//	}
//}

//package main
//
//import "fmt"
//
//func main()  {
//	//指针的指针
//	//var a int
//	//var ptr *int
//	//var pptr *int
//	a := 1234
//	/* 指针 ptr 赋值 */
//	ptr := &a
//	/* 指针 pptr 赋值 */
//	pptr := &ptr
//	fmt.Println(pptr, ptr) // 0xc000006028 0xc00000a0c8
//	/* 获取pptr的值 */
//	fmt.Printf("变量 a = %d \n", a) // 变量 a = 1234
//	fmt.Printf("指针变量 *ptr = %d \n", *ptr) // 指针变量 *ptr = 1234
//	fmt.Printf("指向指针的指针变量 **pptr = %d \n", **pptr) // 指向指针的指针变量 **pptr = 1234
//}

//package main
//
//import "fmt"
//
//func main()  {
//	// 函数传int类型的值与引用的对比
//	a := 10
//	fmt.Printf("1. 变量a的内存地址: %p， 值为：%v\n", &a, a) // 1. 变量a的内存地址: 0xc00000a0c8， 值为：10
//	fmt.Printf("========int型变量a的内存地址：%p\n", a) // ========int型变量a的内存地址：%!p(int=10)
//	changeIntVal(a)
//	fmt.Printf("2. changeIntVal函数调用之后：变量a的内存地址: %p， 值为：%v\n", &a, a) // 2. changeIntVal函数调用之后：变量a的内存地址: 0xc00000a0c8， 值为：10
//	changeIntPtr(&a)
//	fmt.Printf("3. changeIntPtr函数调用之后：变量a的内存地址: %p， 值为：%v\n", &a, a) // 3. changeIntPtr函数调用之后：变量a的内存地址: 0xc00000a0c8， 值为：10
//}
//func changeIntVal(a int) {
//	fmt.Printf("==========changeIntVal函数内：值参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeIntVal函数内：值参数a的内存地址：0xc00000a100，值为：10
//}
//
//func changeIntPtr(a *int) {
//	fmt.Printf("==========changeIntPtr函数内：指针参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeIntPtr函数内：指针参数a的内存地址：0xc000006030，值为：0xc00000a0c8
//}

//package main
//
//import "fmt"
//
//func main() {
//	a := []int{1,2,3,4}
//	fmt.Printf("1. 变量a的内存地址: %p， 值为：%v\n", &a, a) // 1. 变量a的内存地址: 0xc0000044a0， 值为：[1 2 3 4]
//	fmt.Printf("========切片型变量a的内存地址：%p\n", a) // ========切片型变量a的内存地址：0xc00000e360
//	// 传值
//	changeSliceVal(a)
//	fmt.Printf("2. changeSliceVal函数调用之后：变量a的内存地址: %p， 值为：%v\n", &a, a) // 2. changeSliceVal函数调用之后：变量a的内存地址: 0xc0000044a0， 值为：[99 2 3 4]
//	// 传引用
//	changeSlicetr(&a)
//	fmt.Printf("3. changeSlicetr函数调用之后：变量a的内存地址: %p， 值为：%v\n", &a, a) // 3. changeSlicetr函数调用之后：变量a的内存地址: 0xc0000044a0， 值为：[99 100 3 4]
//}
//func changeSliceVal(a []int) {
//	fmt.Printf("==========changeSliceVal函数内：值参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeSliceVal函数内：值参数a的内存地址：0xc000004520，值为：[1 2 3 4]
//	fmt.Printf("==========changeSlicePtr函数内：值参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeSlicePtr函数内：值参数a的内存地址：0xc000004520，值为：[1 2 3 4]
//	a[0] = 99
//}
//
//func changeSlicetr(a *[]int) {
//	fmt.Printf("==========changeSlicetr函数内：指针参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeSlicetr函数内：指针参数a的内存地址：0xc000006030，值为：&[99 2 3 4]
//	(*a)[1] = 100
//}

//package main
//
//import "fmt"
//
//func main() {
//	a := [4]int{1,2,3,4}
//	fmt.Printf("1. 变量a的内存地址: %p， 值为：%v\n", &a, a) // 1. 变量a的内存地址: 0xc00000e360， 值为：[1 2 3 4]
//	fmt.Printf("========数组型变量a的内存地址：%p\n", a) // ========数组型变量a的内存地址：%!p([4]int=[1 2 3 4])
//	// 传值
//	changeArrayVal(a)
//	fmt.Printf("2. changeArrayVal函数调用之后：变量a的内存地址: %p， 值为：%v\n", &a, a) // 2. changeArrayVal函数调用之后：变量a的内存地址: 0xc00000e360， 值为：[1 2 3 4]
//	// 传引用
//	changeArraytr(&a)
//	fmt.Printf("3. changeArraytr函数调用之后：变量a的内存地址: %p， 值为：%v\n", &a, a) // 3. changeArraytr函数调用之后：变量a的内存地址: 0xc00000e360， 值为：[1 100 3 4]
//}
//func changeArrayVal(a [4]int) {
//	fmt.Printf("==========changeArrayVal函数内：值参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeArrayVal函数内：值参数a的内存地址：0xc00000e3e0，值为：[1 2 3 4]
//	fmt.Printf("==========changeArrayPtr函数内：值参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeArrayPtr函数内：值参数a的内存地址：0xc00000e3e0，值为：[1 2 3 4]
//	a[0] = 99
//}
//
//func changeArraytr(a *[4]int) {
//	fmt.Printf("==========changeArraytr函数内：指针参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeArraytr函数内：指针参数a的内存地址：0xc000006030，值为：&[1 2 3 4]
//	(*a)[1] = 100
//}

//package main
//
//import "fmt"
//
//type Teacher struct {
//	name string
//	age int
//	married bool
//	sex int8
//}
//
//func main() {
//	a := Teacher{"Steven", 35, true, 1}
//	fmt.Printf("1. 变量a的内存地址: %p， 值为：%v\n", &a, a) // 1. 变量a的内存地址: 0xc000066440， 值为：{Steven 35 true 1}
//	fmt.Printf("struct型变量a的内存地址：%p\n", a) // struct型变量a的内存地址：%!p(main.Teacher={Steven 35 true 1})
//	// 传值
//	changeStructVal(a)
//	fmt.Printf("2. changeStructVal函数调用之后：变量a的内存地址: %p， 值为：%v\n", &a, a) // 2. 2. changeStructVal函数调用之后：变量a的内存地址: 0xc000066440， 值为：{Steven 35 true 1}
//	// 传引用
//	changeStructPtr(&a)
//	fmt.Printf("3. changeStructPtr函数调用之后：变量a的内存地址: %p， 值为：%v\n", &a, a) // 3. changeStructPtr函数调用之后：变量a的内存地址: 0xc000066440， 值为：{Daniel 33 false 1}
//}
//func changeStructVal(a Teacher) {
//	fmt.Printf("==========changeStructVal函数内：值参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeStructVal函数内：值参数a的内存地址：0xc0000664c0，值为：{Steven 35 true 1}
//	fmt.Printf("==========changeStructPtr函数内：值参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeStructPtr函数内：值参数a的内存地址：0xc0000664c0，值为：{Steven 35 true 1}
//	a.name = "josh"
//	a.age = 29
//	a.married =false
//}
//
//func changeStructPtr(a *Teacher) {
//	fmt.Printf("==========changeStructPtr函数内：指针参数a的内存地址：%p，值为：%v\n", &a, a) // ==========changeStructPtr函数内：指针参数a的内存地址：0xc00009a020，值为：&{Steven 35 true 1}
//	(*a).name = "Daniel"
//	(*a).age = 33
//	(*a).married = false
//}

//package main
//
//import "fmt"
//
//func main()  {
//	a := [4]float64{67.12, 89.8, 21, 34}
//	b := [...]int{1,2,3,4,5}
//	fmt.Printf("数组a的长度为 %d, 数组 b的长度为 %d\n", len(a), len(b)) // 数组a的长度为 4, 数组 b的长度为 5
//}

//package main
//
//import "fmt"
//
//func main()  {
//	a := [4]float64{67.12, 89.8, 21, 34}
//	b := [...]int{1,2,3,4,5}
//	// 遍历数组方式1		67.12	89.8	21	34
//	for i := 0; i < len(a); i++ {
//		fmt.Print(a[i], "\t")
//	}
//	fmt.Println()
//	// 遍历数组方式2		1	2	3	4	5
//	for _, value := range b{
//		fmt.Print(value, "\t")
//	}
//}

//package main
//
//import "fmt"
//
//func main()  {
//	var numbers  =  make([]int, 3, 5)
//	fmt.Printf("%T \n", numbers)  // []int
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)// len=3 cap=5 slice=[0 0 0]
//}

//package main
//
//import "fmt"
//
//func main()  {
//	/*创建切片*/
//	numbers := []int{0,1,2,3,4,5,6,7,8,9}
//	printSlice(numbers) // len=10 cap=10 slice=[0 1 2 3 4 5 6 7 8 9]
//	/*打印原始切片*/
//	fmt.Println("numbers = ", numbers) // numbers =  [0 1 2 3 4 5 6 7 8 9]
//	/*打印切片从索引1（包含）到索引4（不包含）*/
//	fmt.Println("numbers[1:4] =", numbers[1:4]) // numbers[1:4] = [1 2 3]
//	/*默认下限为0*/
//	fmt.Println("numbers[:3] = ", numbers[:3]) // numbers[:3] =  [0 1 2]
//	/*默认上限为len(s)*/
//	fmt.Println("numbers[4:] = ", numbers[4:]) // numbers[4:] =  [4 5 6 7 8 9]
//	/*打印子切片从索引0（包含）到索引2（不包含）*/
//	number2 := numbers[:2]
//	printSlice(number2) // len=2 cap=10 slice=[0 1]
//	/*打印子切片从索引2（包含）到索引6（不包含）*/
//	number3 := numbers[2:5]
//	printSlice(number3) // len=3 cap=8 slice=[2 3 4]
//}
//
//func printSlice(x []int)  {
//	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//}

//package main
//
//import "fmt"
//
//func main()  {
//	sliceCap()
//}
//func sliceCap()  {
//	arr0 := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
//	fmt.Println("cap(arr0) = ", cap(arr0), arr0) // cap(arr0) =  11 [a b c d e f g h i j k]
//	//截取数组，形成切片
//	s01 := arr0[2:8]
//	fmt.Printf("%T \n", s01) // []string
//	fmt.Println("cap(s01) = ", cap(s01), s01) // cap(s01) =  9 [c d e f g h]
//	s02 := arr0[4:7]
//	fmt.Println("cap(s02) = ", cap(s02), s02) //cap(s02) =  7 [e f g]
//	// 截取切片， 形成切片
//	s03 := s01[3:9]
//	fmt.Println("截取s01[3:9]后形成s03: ", s03) // 截取s01[3:9]后形成s03:  [f g h i j k]
//	s04 := s01[4:7]
//	fmt.Println("截取s01[4:7]后形成s03: ", s04) // 截取s01[4:7]后形成s03:  [g h i]
//	//切片是引用类型
//	s04[0] = "x"
//	fmt.Print(arr0, s01, s02, s03, s04) // [a b c d e f x h i j k] [c d e f x h] [e f x] [f x h i j k] [x h i]
//}

//package main
//
//import "fmt"
//
//func main()  {
//	fmt.Println("1. ============================")
//	numbers := make([]int, 0, 20)
//	PrintSlices("numbers:", numbers)
//	/* 向切片添加一个元素 */
//	numbers = append(numbers, 1)
//	PrintSlices("numbers:", numbers)
//	/* 同时添加多个元素 */
//	numbers = append(numbers, 2, 3, 4, 5, 6, 7)
//	PrintSlices("numbers:", numbers)
//	fmt.Println("2. ============================")
//	// 追加一个切片
//	s1 := []int{100, 200, 300, 400, 500, 600, 700}
//	numbers = append(numbers, s1...)
//	PrintSlices("numbers:", numbers)
//	fmt.Println("3. ============================")
//	// 切片删除元素
//	// 删除第一个元素
//	numbers = numbers[1:]
//	PrintSlices("numbers:", numbers)
//	//删除最后一个元素
//	numbers = numbers[:len(numbers)-1]
//	PrintSlices("numbers:", numbers)
//	//删除中间一个元素
//	a := int(len(numbers)/2)
//	fmt.Println("中间数：",a)
//	numbers = append(numbers[:a], numbers[a:]...)
//	PrintSlices("numbers:", numbers)
//	fmt.Println("4. ============================")
//	/*创建切片numbers1是之前切片的两倍容量*/
//	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
//	/* 复制numbers的内容到numbers1 */
//	count := copy(numbers1, numbers)
//	fmt.Println("复制个数：", count)
//	PrintSlices("numbers1", numbers1)
//	numbers[len(numbers)-1] = 99
//	numbers1[0] = 100
//	/* Numbers1与numbers两者不存在联系，numbers发生变化时，
//	numbers1是不会随着变化的，也就是说copy函数不会建立两个切片的联系
//	 */
//	PrintSlices("numbers:", numbers)
//	PrintSlices("numbers1:", numbers1)
//}
//// 输出切片格式化信息
//func PrintSlices(name string, x[]int)  {
//	fmt.Print(name, "\t")
//	fmt.Printf("addr:%p \t len=%d \t cap=%d \t slice=%v\n", x, len(x), cap(x), x)
//}
///*
//输出结果：
//
//1. ============================
//numbers:	addr:0xc000092000 	 len=0 	 cap=20 	 slice=[]
//numbers:	addr:0xc000092000 	 len=1 	 cap=20 	 slice=[1]
//numbers:	addr:0xc000092000 	 len=7 	 cap=20 	 slice=[1 2 3 4 5 6 7]
//2. ============================
//numbers:	addr:0xc000092000 	 len=14 	 cap=20 	 slice=[1 2 3 4 5 6 7 100 200 300 400 500 600 700]
//3. ============================
//numbers:	addr:0xc000092008 	 len=13 	 cap=19 	 slice=[2 3 4 5 6 7 100 200 300 400 500 600 700]
//numbers:	addr:0xc000092008 	 len=12 	 cap=19 	 slice=[2 3 4 5 6 7 100 200 300 400 500 600]
//中间数： 6
//numbers:	addr:0xc000092008 	 len=12 	 cap=19 	 slice=[2 3 4 5 6 7 100 200 300 400 500 600]
//4. ============================
//复制个数： 12
//numbers1	addr:0xc000096000 	 len=12 	 cap=38 	 slice=[2 3 4 5 6 7 100 200 300 400 500 600]
//numbers:	addr:0xc000092008 	 len=12 	 cap=19 	 slice=[2 3 4 5 6 7 100 200 300 400 500 99]
//numbers1:	addr:0xc000096000 	 len=12 	 cap=38 	 slice=[100 3 4 5 6 7 100 200 300 400 500 600]
// */

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func main()  {
//	// 使用那种方式初始化切片更加高效？
//	var sa [] string
//	// sa := make([]string, 0, 20)
//	printSliceMsg(sa)
//	// 当使用append追加元素到切片时，若容量不够，go就会创建一个新的切片变量来存储元素
//	for i := 0; i < 15; i++{
//		sa = append(sa, strconv.Itoa(i))
//		printSliceMsg(sa)
//	}
//	printSliceMsg(sa)
//}
//func printSliceMsg(sa []string){
//	fmt.Printf("addr:%p\t len:%v\tcap:%d\tvalue:%v\n", sa, len(sa), cap(sa), sa)
//}
///*
//运行结果：
//addr:0x0	 len:0	cap:0	value:[]
//addr:0xc0000581e0	 len:1	cap:1	value:[0]
//addr:0xc0000664a0	 len:2	cap:2	value:[0 1]
//addr:0xc0000a2000	 len:3	cap:4	value:[0 1 2]
//addr:0xc0000a2000	 len:4	cap:4	value:[0 1 2 3]
//addr:0xc0000a4000	 len:5	cap:8	value:[0 1 2 3 4]
//addr:0xc0000a4000	 len:6	cap:8	value:[0 1 2 3 4 5]
//addr:0xc0000a4000	 len:7	cap:8	value:[0 1 2 3 4 5 6]
//addr:0xc0000a4000	 len:8	cap:8	value:[0 1 2 3 4 5 6 7]
//addr:0xc0000a6000	 len:9	cap:16	value:[0 1 2 3 4 5 6 7 8]
//addr:0xc0000a6000	 len:10	cap:16	value:[0 1 2 3 4 5 6 7 8 9]
//addr:0xc0000a6000	 len:11	cap:16	value:[0 1 2 3 4 5 6 7 8 9 10]
//addr:0xc0000a6000	 len:12	cap:16	value:[0 1 2 3 4 5 6 7 8 9 10 11]
//addr:0xc0000a6000	 len:13	cap:16	value:[0 1 2 3 4 5 6 7 8 9 10 11 12]
//addr:0xc0000a6000	 len:14	cap:16	value:[0 1 2 3 4 5 6 7 8 9 10 11 12 13]
//addr:0xc0000a6000	 len:15	cap:16	value:[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
//addr:0xc0000a6000	 len:15	cap:16	value:[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
//*/

//package main
//
//import "fmt"
//
//func main() {
//	// 1. 声明时同时初始化
//	var country = map[string] string{
//		"China": "Beijing",
//		"Japan": "Tokyo",
//		"India": "New Delhi",
//		"France": "Paris",
//		"Italy": "Rome",
//	}
//	fmt.Println(country)
//	//短变量声明初始化
//	rating := map[string]float64{"C":5, "Go":4.5, "Python": 4.5, "C++": 3}
//	fmt.Println(rating)
//	//2. 创建map后赋值
//	countryMap := make(map[string]string)
//	countryMap["China"] = "Beijing"
//	countryMap["Japan"] = "Tokyo"
//	countryMap["India"] = "New Delhi"
//	countryMap["France"] = "Paris"
//	countryMap["Italy"] = "Rome"
//	//3. 遍历map
//	//(1) key, value都遍历
//	for k,v := range countryMap{
//		fmt.Println("国家", k, "首都", v)
//	}
//	fmt.Println("======================")
//	//(2)只展示value
//	for _,v := range countryMap{
//		fmt.Println("国家", "首都", v)
//	}
//	fmt.Println("======================")
//	//(3)只展示key
//	for k,_ := range countryMap{
//		fmt.Println("国家", k, "首都")
//	}
//}
///*
//运行结果：
//map[China:Beijing France:Paris India:New Delhi Italy:Rome Japan:Tokyo]
//map[C:5 C++:3 Go:4.5 Python:4.5]
//国家 Japan 首都 Tokyo
//国家 India 首都 New Delhi
//国家 France 首都 Paris
//国家 Italy 首都 Rome
//国家 China 首都 Beijing
//======================
//国家 首都 Beijing
//国家 首都 Tokyo
//国家 首都 New Delhi
//国家 首都 Paris
//国家 首都 Rome
//======================
//国家 China 首都
//国家 Japan 首都
//国家 India 首都
//国家 France 首都
//国家 Italy 首都
//
// */

//package main
//
//import "fmt"
//
//func main()  {
//	countryMap := make(map[string]string)
//	countryMap["China"] = "Beijing"
//	countryMap["Japan"] = "Tokyo"
//	countryMap["India"] = "New Delhi"
//	countryMap["France"] = "Paris"
//	countryMap["Italy"] = "Rome"
//	//查看元素是否在map中
//	value, ok := countryMap["England"]
//	fmt.Printf("%q\n", value)
//	fmt.Printf("%T, %v\n", ok, ok)
//	if ok{
//		fmt.Printf("首都:", value)
//	}else {
//		fmt.Println("首都信息为检索到！")
//	}
//	//或者
//	if value, ok := countryMap["USA"];ok{
//		fmt.Printf("首都:", value)
//	}else {
//		fmt.Println("首都信息为检索到！")
//	}
//}
///*
//运行结果：
//""
//bool, false
//首都信息为检索到！
//首都信息为检索到！
// */

//package main
//
//import "fmt"
//
//func main()  {
//	countryMap := make(map[string]string)
//	countryMap["China"] = "Beijing"
//	countryMap["Japan"] = "Tokyo"
//	countryMap["India"] = "New Delhi"
//	countryMap["France"] = "Paris"
//	countryMap["Italy"] = "Rome"
//	//根据key删除map中的某个元素
//	fmt.Println("删除前:", countryMap)
//	if _,ok := countryMap["Italy"]; ok{
//		delete(countryMap, "Italy")
//	}
//	fmt.Println("删除后:", countryMap)
//	// 清空map
//	//countryMap = map[string]string{}
//	countryMap = make(map[string]string)
//	fmt.Println("清空后：",countryMap)
//}
///*
//运行结果：
//删除前: map[China:Beijing France:Paris India:New Delhi Italy:Rome Japan:Tokyo]
//删除后: map[China:Beijing France:Paris India:New Delhi Japan:Tokyo]
//清空后： map[]
// */

//
//package main
//
//import (
//	"fmt"
//	"unicode/utf8"
//)
//
//func main()  {
//	s := "我爱Go语言"
//	fmt.Println("字节长度", len(s))
//	fmt.Println("------------------")
//	//for ... range遍历字符串
//	len := 0
//	for i, ch :=range s{
//		fmt.Printf("%d : %x \t", i, ch)
//		len++
//	}
//	fmt.Println("\n字符串长度", len)
//	fmt.Println("------------------")
//	// 遍历所有字节
//	for i, ch :=range []byte(s) {
//		fmt.Printf("%d : %x \t", i, ch)
//	}
//	fmt.Println()
//	fmt.Println("------------------")
//	//遍历所有字符
//	count := 0
//	for i, ch :=range []rune(s) {
//		fmt.Printf("%d : %c \t", i, ch)
//		count++
//	}
//	fmt.Println()
//	fmt.Println("字符串长度", count)
//	fmt.Println("字符串长度", utf8.RuneCountInString(s))
//}
///*
//运行结果：
//字节长度 14
//------------------
//0 : 6211 	3 : 7231 	6 : 47 	7 : 6f 	8 : 8bed 	11 : 8a00
//字符串长度 6
//------------------
//0 : e6 	1 : 88 	2 : 91 	3 : e7 	4 : 88 	5 : b1 	6 : 47 	7 : 6f 	8 : e8 	9 : af 	10 : ad 	11 : e8 	12 : a8 	13 : 80
//------------------
//0 : 我 	1 : 爱 	2 : G 	3 : o 	4 : 语 	5 : 言
//字符串长度 6
//字符串长度 6
//*/

//package main
//
//import (
//	"fmt"
//	"strings"
//	"unicode"
//)
//
//func main()  {
//	// 判断是否包含字串
//	fmt.Println(strings.Contains("seafood", "foo")) // true
//	fmt.Println(strings.Contains("seafood", "bar")) // false
//	fmt.Println(strings.Contains("seafood", "")) // true
//	fmt.Println(strings.Contains("", "")) // true
//	fmt.Println(strings.Contains("steven王2008", "王")) // true
//
//	// 判断字符串是否包含另一字符串中的任一字符
//	fmt.Println(strings.ContainsAny("team", "i")) // false
//	fmt.Println(strings.ContainsAny("failure", "u & i")) // true
//	fmt.Println(strings.ContainsAny("foo", "")) // false
//	fmt.Println(strings.ContainsAny("", "")) // false
//
//	// 判断字符串是否包含unicode码值
//	fmt.Println(strings.ContainsRune("一丁", '丁')) // true
//	fmt.Println(strings.ContainsRune("一丁", 19969)) // true
//
//	// 返回字符串包含另一个字符串的个数
//	fmt.Println(strings.Count("cheese", "e")) // 3
//	fmt.Println(strings.Count("one", "")) // 4
//
//	// 判断字符串是否有前缀字符串
//	fmt.Println(strings.HasPrefix("1000phone news", "1000")) // true
//	fmt.Println(strings.HasPrefix("1000phone news", "1000a")) // false
//
//	// 判断字符串是否有后缀字符串
//	fmt.Println(strings.HasSuffix("1000phone news", "news")) // true
//	fmt.Println(strings.HasSuffix("1000phone news", "new")) // false
//
//	// 返回字符串中另一字符串首次出现的位置
//	fmt.Println(strings.Index("chicken", "ken")) // 4
//	fmt.Println(strings.Index("chicken", "dmr")) // -1
//
//	// 返回字符串中的任一unicode码值首次出现的位置
//	fmt.Println(strings.IndexAny("abcABC120", "教育基地A")) // 3
//
//	// 返回字符串中字符串首次出现的位置
//	fmt.Println(strings.IndexByte("123abc", 'a')) // 3
//
//	// 判断字符串是否包含unicode码值
//	fmt.Println(strings.IndexRune("abcABC120", 'C')) // 5
//	fmt.Println(strings.IndexRune("It培训教育", '教')) // 8
//
//	// 返回字符串中满足函数f(r)==true字符首次出现的位置
//	f := func(c rune) bool{
//		return unicode.Is(unicode.Han, c)
//	}
//	fmt.Println(strings.IndexFunc("Hello123,中国", f)) // 9
//
//	// 返回字符串中字串最后一次出现的位置
//	fmt.Println(strings.LastIndex("Steven learn english", "e")) // 13
//	fmt.Println(strings.Index("go gopher", "go")) // 0
//	fmt.Println(strings.LastIndex("go gopher", "go")) // 3
//	fmt.Println(strings.LastIndex("go gopher", "rodent")) // -1
//
//	// 返回字符串中任意一个unicode码值最后一次出现的位置
//	fmt.Println(strings.LastIndexAny("chicken", "aeiouy")) // 5
//	fmt.Println(strings.LastIndexAny("crwth", "aeiouy")) // -1
//
//	// 返回字符中字符最后一次出现得位置
//	fmt.Println(strings.LastIndexByte("abcABCA123", 'A')) // 6
//
//	// 返回字符串中满足函数f(r)==true字符最后一次出现的位置
//	f1 := func(c rune) bool{
//		return unicode.Is(unicode.Han, c)
//	}
//	fmt.Println(strings.LastIndexFunc("Hello,世界", f1)) // 9
//	fmt.Println(strings.LastIndexFunc("Hello word,中国人", f1)) // 17
//
//	res := GetFileSuffix("abc.xyz.lmn.jpg")
//	fmt.Println(res) // jpg
//}
//// 获取文件后缀
//func GetFileSuffix(str string) string{
//	arr := strings.Split(str, ".")
//	return arr[len(arr)-1]
//}

//package main
//
//import (
//	"fmt"
//	"strings"
//	"unicode"
//)
//
//func main(){
//	// 将字符串以空白字符分割，并返回一个切片
//	fmt.Println(strings.Fields("  abc 123 ABC xyz XYZ")) // [abc 123 ABC xyz XYZ]
//	fmt.Println(strings.Fields("abcdefghijklmnopqrstuvwxyz")) // [abcdefghijklmnopqrstuvwxyz]
//
//	// 将字符串以满足在f(r)==true的字符进行分割，返回一个切片
//	f := func(c rune) bool {
//		//return c == '='
//		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
//	}
//	fmt.Println(strings.FieldsFunc("abc@123*ABC&xyz%XYZ", f)) // [abc 123 ABC xyz XYZ]
//
//	// 将字符串以sep作为分隔符进行分割，分割后字符最后去掉sep
//	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) // ["a" "b" "c"]
//	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
//	fmt.Printf("%q\n", strings.Split(" xyz ", "")) // [" " "x" "y" "z" " "]
//	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins")) // [""]
//
//	// 将字符串s以sep作为分隔符进行分割，分割后字符最后附上sep，n决定返回的切片数
//	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) // ["a" "b,c"]
//	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 1)) // ["a,b,c"]
//
//	// 将字符串s以sep作为分隔符进行分割，分割后的字符最后附上sep
//	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]
//
//	// 将字符串s以sep作为分隔符进行分割，分割后字符最后附上sep，n决定返回的切片数
//	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]
//}

//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main()  {
//	// 将字符串s每个单词大写返回
//	fmt.Println(strings.Title("her royal highness")) // Her Royal Highness
//
//	// 将字符串s转换成大写返回
//	fmt.Println(strings.ToTitle("louD noises")) // LOUD NOISES
//
//	// 将字符串s转化成小写返回
//	fmt.Println(strings.ToLower("GOPHER")) // gopher
//
//	// 将字符串s转换成大写返回
//	fmt.Println(strings.ToUpper("gopher")) // GOPHER
//}

//package  main
//
//import (
//	"fmt"
//	"strings"
//	"unicode"
//)
//
//func main()  {
//	// 将字符串s首尾包含在cutset中的任一字符去掉返回
//	fmt.Println(strings.Trim("  steven wang    ", " ")) // steven wang
//
//	// 将字符串s首尾满足函数f(r)==true的字符去掉返回
//	f1 := func(c rune) bool {
//		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
//	}
//	fmt.Println(strings.TrimFunc("！@#￥%steven wang%￥#@", f1)) // steven wang
//
//	// 将字符串s左边包含在cutset中的任一字符去掉返回
//	fmt.Println(strings.Trim("  steven wang    m ", " ")) // steven wang    m
//
//	// 将字符串s左边满足函数f(r)==true的字符去掉返回
//	f2 := func(c rune) bool {
//		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
//	}
//	fmt.Println(strings.TrimFunc("！@#￥%steven wang%￥#@", f2)) // steven wang
//
//	// 将字符串s右边包含在cutset中的任一字符去掉返回
//	fmt.Println(strings.Trim("  steven wang    ", " ")) //steven wang
//
//	// 将字符串s右边满足函数f(r)==true的字符去掉返回
//	f3 := func(c rune) bool {
//		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
//	}
//	fmt.Println(strings.TrimFunc("！@#￥%steven wang%￥#@", f3)) //steven wang
//
//	// 将字符串s首尾空白去掉返回
//	fmt.Println(strings.TrimSpace(" \t\n a lone goper \n\t\r\n")) // a lone goper
//
//	// 将字符串s中 前缀字符串prefix去掉返回
//	fmt.Println(strings.TrimPrefix("Goodbye,world", "Goodbye")) // ,world
//
//	// 将字符串s中后缀字符串suffix去掉返回
//	fmt.Println(strings.TrimSuffix("Hello, goodbye, etc!", "goodbye, etc!"))
//}

//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main()  {
//	// 按字典顺序比较a和b字符串大小
//	fmt.Println(strings.Compare("abc", "bcd")) // -1
//	fmt.Println("abc" < "bcd") // true
//
//	// 判断s和t两个UTF-8字符串是否相等，忽略大小写
//	fmt.Println(strings.EqualFold("Go", "go")) // true
//
//	// 将字符串s重复count次返回
//	fmt.Println("g" + strings.Repeat("o", 8) + "le") // goooooooole
//
//	// 替换字符串s中old字符为new字符并返回，n < 0时替换所有old字符串
//	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", 2)) // 张老大 张老二 王老三
//	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", -1)) // 张老大 张老二 张老三
//	fmt.Println(strings.ReplaceAll("王老大 王老二 王老三", "王", "张")) // 张老大 张老二 张老三
//
//	// 将a中的所有字符连接成一个字符串，使用字符串sep作为分隔符
//	s := []string{"abc", "ABC", "123"}
//	fmt.Println(strings.Join(s, ",")) // abc,ABC,123
//	fmt.Println(strings.Join(s, " ")) // abc ABC 123
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func main()  {
//	// 将字符串类型转换为int类型
//	a, _ := strconv.Atoi("100")
//	fmt.Printf("%T, %v\n", a, a+2) // int, 102
//	fmt.Println("-------------------------")
//
//	//解释给定基数的字符串s并返回相应的值i
//	num, _ := strconv.ParseInt("-4e00", 16, 64) // int64, -19968
//	fmt.Printf("%T, %v\n", num, num)
//	num, _ = strconv.ParseInt("01100001", 2, 64) // int64, 97
//	fmt.Printf("%T, %v\n", num, num)
//	num, _ = strconv.ParseInt("-01100001", 10, 64) // int64, -1100001
//	fmt.Printf("%T, %v\n", num, num)
//	num, _ = strconv.ParseInt("4e00", 10, 64) // int64, 0
//	fmt.Printf("%T, %v\n", num, num)
//	fmt.Println("-------------------------")
//
//	// ParseUint类似ParseInt，但是用于无符号数字
//	num1, _ := strconv.ParseUint("-4e00", 16, 64) // uint64, 0
//	fmt.Printf("%T, %v\n", num1, num1)
//	num1, _ = strconv.ParseUint("01100001", 2, 64) // uint64, 97
//	fmt.Printf("%T, %v\n", num1, num1)
//	num1, _ = strconv.ParseUint("-01100001", 10, 64) // uint64, 0
//	fmt.Printf("%T, %v\n", num1, num1)
//	num1, _ = strconv.ParseUint("4e00", 10, 64) // uint64, 0
//	fmt.Printf("%T, %v\n", num1, num1)
//	fmt.Println("-------------------------")
//
//	// ParseFloat将字符串s转换为float类型
//	pi := "3.1415926"
//	num2, _ := strconv.ParseFloat(pi, 64)
//	fmt.Printf("%T, %v\n", num2, num2*2) // float64, 6.2831852
//	fmt.Println("-------------------------")
//
//	// 将字符串转换为Bool类型
//	flag, _ := strconv.ParseBool("steven")
//	fmt.Printf("%T, %v\n", flag, flag) // bool, false
//	fmt.Println("-------------------------")
//}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//
//func main()  {
//	// Int转换成string
//	s1 := strconv.Itoa(199)
//	fmt.Printf("%T, %v, 长度：%d \n", s1, s1, len(s1)) // string, 199, 长度：3
//
//	// 返回给定基数的i的字符串表示
//	s2 := strconv.FormatInt(-19968, 16)
//	fmt.Printf("%T, %v, 长度：%d \n", s2, s2, len(s2)) // string, -4e00, 长度：5
//	s2 = strconv.FormatInt(-40869, 16)
//	fmt.Printf("%T, %v, 长度：%d \n", s2, s2, len(s2)) // string, -9fa5, 长度：5
//	fmt.Println("--------------------------")
//	s3 := strconv.FormatInt(19968, 16)
//	fmt.Printf("%T, %v, 长度：%d \n", s3, s3, len(s3)) // string, 4e00, 长度：4
//	s2 = strconv.FormatInt(40869, 16)
//	fmt.Printf("%T, %v, 长度：%d \n", s3, s3, len(s3)) // string, 4e00, 长度：4
//
//	// 将浮点数f转换为字符串
//	s4 := strconv.FormatFloat(3.1415926, 'g', -1, 64)
//	fmt.Printf("%T, %v, 长度：%d \n", s4, s4, len(s4)) // string, 3.1415926, 长度：9
//}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main()  {
//	time1 := time.Now()
//	TestTime()
//	time2 := time.Now()
//	fmt.Println(time2.Sub(time1).Seconds())
//}
//func TestTime(){
//	t := time.Now()
//	fmt.Println("1. ", t) // 1.  2020-04-19 01:56:26.5039714 +0800 CST m=+0.004024301
//	fmt.Println("2. ", t.Local()) // 2.  2020-04-19 01:56:26.5039714 +0800 CST
//	fmt.Println("3. ", t.UTC()) // 3.  2020-04-18 17:56:26.5039714 +0000 UTC
//	t = time.Date(2020, time.January, 1, 1, 1, 1, 0, time.Local)
//	fmt.Printf("4. 本地时间%s , 国际统一时间：%s \n", t, t.UTC())  // 4. 本地时间2020-01-01 01:01:01 +0800 CST , 国际统一时间：2019-12-31 17:01:01 +0000 UTC
//	t, _ = time.Parse("2006-01-02 15:04:23", "2018-07-19 04:23:34")
//	fmt.Println("5. ", t) // 5.  0001-01-01 00:00:00 +0000 UTC
//	fmt.Println("6. ", time.Now().Format("2006-01-02 15:04:23")) // 6.  2020-04-19 01:56:191
//	fmt.Println("7. ", time.Now().String()) // 7.  2020-04-19 01:56:26.5149209 +0800 CST m=+0.014973801
//	fmt.Println("8. ", time.Now().Unix()) // 8.  1587232586
//	fmt.Println("9. ", time.Now().UnixNano()) // 9.  1587232586514920900
//	fmt.Println("10. ", t.Equal(time.Now())) // 10.  false
//	fmt.Println("11. ", t.Before(time.Now())) // 11.  true
//	fmt.Println("12. ", t.After(time.Now())) // 12.  false
//	year, month, day := time.Now().Date()
//	fmt.Println("13. ", year, month, day) // 13.  2020 April 19
//	fmt.Println("14. ", time.Now().Year()) // 14.  2020
//	fmt.Println("15. ", time.Now().Month()) // 15.  April
//	fmt.Println("16. ", time.Now().Day()) // 16.  19
//	fmt.Println("17. ", time.Now().Weekday()) // 17.  Sunday
//	hour, minute, second := time.Now().Clock()
//	fmt.Println("18. ", hour, minute, second) // 18.  1 56 26
//	fmt.Println("19. ", time.Now().Hour()) // 19.  1
//	fmt.Println("20. ", time.Now().Minute()) // 20.  56
//	fmt.Println("21. ", time.Now().Second()) // 21.  26
//	fmt.Println("22. ", time.Now().Nanosecond()) // 22.  514920900
//	fmt.Println("23. ", time.Now().Sub(time.Now())) // 23.  0s
//	fmt.Println("24. ", time.Now().Sub(time.Now()).Hours()) // 24.  0
//	fmt.Println("25. ", time.Now().Sub(time.Now()).Minutes()) // 25.  0
//	fmt.Println("26. ", time.Now().Sub(time.Now()).Seconds()) // 26.  0
//	fmt.Println("27. ", time.Now().Sub(time.Now()).Nanoseconds()) // 27.  0
//	fmt.Println("28. ", "时间间距： ", t.Sub(time.Now()).String()) // 28.  时间间距：  -2562047h47m16.854775808s
//	d, _ := time.ParseDuration("1h30m")
//	fmt.Println("29. ", d) // 29.  1h30m0s
//	fmt.Println("30. ", "交卷时间：", time.Now().Add(d)) // 30.  交卷时间： 2020-04-19 03:26:26.5149209 +0800 CST m=+5400.014973801
//	fmt.Println("31. ", "一年一个月零一天之后的日期：", time.Now().AddDate(1,1,1)) // 31.  一年一个月零一天之后的日期： 2021-05-20 01:56:26.5149209 +0800 CST
//}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main()  {
//	fmt.Println(math.IsNaN(3.4)) // false
//	fmt.Println(math.Ceil(1.000001)) // 2
//	fmt.Println(math.Floor(1.999999)) // 1
//	fmt.Println(math.Trunc(1.999999)) // 1
//	fmt.Println(math.Abs(-1.3)) // 1.3
//	fmt.Println(math.Max(-1.3, 0)) // 0
//	fmt.Println(math.Min(-1.3, 0)) // -1.3
//	fmt.Println(math.Dim(-12, -19)) // 7
//	fmt.Println(math.Dim(-12, 19)) // 0
//	fmt.Println(math.Mod(9, 4)) // 1
//	fmt.Println(math.Sqrt(9)) // 3
//	fmt.Println(math.Cbrt(8)) // 2
//	fmt.Println(math.Hypot(3, 4)) // 5
//	fmt.Println(math.Pow(2, 8)) // 256
//	fmt.Println(math.Log(1)) // 0
//	fmt.Println(math.Log2(16)) // 4
//	fmt.Println(math.Log10(100)) // 2
//}

/*	输入输出	*/
//package main
//
//import "fmt"
//
//func main()  {
//	username := ""
//	age := 0
//	fmt.Scanln(&username, &age)
//	fmt.Println("账号信息为：", username, age)
//}

///*	结构体	*/
//package main
//
//import "fmt"
//
//// 定义Teacher结构体
//type Teacher struct {
//	name string
//	age int
//	sex byte
//}
//
//func main()  {
//	// 1. var声明方式实例化结构体，初始化方式为：对象.属性=值
//	var t1 Teacher
//	fmt.Println(t1) // { 0 0}
//	fmt.Printf("t1:%T, %v, %q\n", t1, t1, t1) // t1:main.Teacher, { 0 0}, {"" '\x00' '\x00'}
//	t1.name = "Steven"
//	t1.age = 35
//	t1.sex = 1
//	fmt.Println(t1) // {Steven 35 1}
//	fmt.Println("------------------------------------")
//	// 2.变量简短声明格式实例化结构体，初始化方式为：对象.属性=值
//	t2 := Teacher{}
//	t2.name = "David"
//	t2.age = 30
//	t2.sex = 1
//	fmt.Println(t2) // {David 30 1}
//	fmt.Println("------------------------------------")
//	// 3. 变量简短声明格式实例化结构体，声明时初始化，初始化方式为：属性：值， 属性：值可以同行，也可以换行(类似于map的用法）
//	t3 := Teacher{
//		name: "Josh",
//		age:  28,
//		sex:  1,
//	}
//	fmt.Println(t3) // {Josh 28 1}
//	t3 = Teacher{name:"Josh2", age:27, sex:1}
//	fmt.Println(t3) // {Josh2 27 1}
//	fmt.Println("------------------------------------")
//	// 4. 变量简短声明格式实例化结构体，声明时初始化，不屑属性明，按照属性顺序只写属性值
//	t4 := Teacher{"Ruby", 30, 0}
//	fmt.Println(t4) // {Ruby 30 0}
//	fmt.Println("------------------------------------")
//}

// package main

// import "fmt"

// // 定义结构体emp
// type Emp struct {
// 	name string
// 	age  int8
// 	sex  byte
// }

// func main() {
// 	// 使用new()内置函数实例化struct
// 	emp1 := new(Emp)
// 	fmt.Printf("emp: %T, %v, %p\n", emp1, emp1, emp1) // emp: *main.Emp, &{ 0 0}, 0xc0000044a0
// 	(*emp1).name = "David"
// 	(*emp1).age = 30
// 	(*emp1).sex = 1
// 	fmt.Println(emp1)  // &{David 30 1}
// 	fmt.Println(*emp1) // {David 30 1}
// 	// 语法糖写法
// 	emp1.name = "David2"
// 	emp1.age = 31
// 	emp1.sex = 1
// 	fmt.Println(emp1)  // &{David2 31 1}
// 	fmt.Println(*emp1) // {David2 31 1}
// 	fmt.Println("-----------------------")
// 	SyntacticSugar()
// }
// func SyntacticSugar() {
// 	// 数组中的语法糖
// 	arr := [4]int{10, 20, 30, 40}
// 	arr2 := &arr
// 	fmt.Println((*arr2)[len(arr)-1]) // 40
// 	fmt.Println(arr2[0])             // 10
// 	// 切片中的语法糖
// 	arr3 := []int{100, 200, 300, 400}
// 	arr4 := &arr3
// 	fmt.Println((*arr4)[len(arr3)-1]) // 400
// }

package main

import "fmt"

// Human is struct, description a person
type Human struct {
	name string
	age  int8
	sex  byte
}

func main() {
	// 1.初始化Human
	h1 := Human{"Steven", 35, 1}
	fmt.Printf("h1: %T, %v, %p \n", h1, h1, &h1) // h1: main.Human, {Steven 35 1}, 0xc0000044a0
	fmt.Println("-----------------------------")
	// 2.复制结构体对象
	h2 := h1
	h2.name = "David"
	h2.age = 30
	fmt.Printf("h2修改后：%T, %v, %p \n", h2, h2, &h2) // h2修改后：main.Human, {David 30 1}, 0xc000004520
	fmt.Printf("h1: %T, %v, %p \n", h1, h1, &h1)   // h1: main.Human, {Steven 35 1}, 0xc0000044a0
	fmt.Println("-----------------------------")
	// 3.将结构体对象作为参数传递
	changeName(h1)
	fmt.Printf("h1: %T, %v, %p \n", h1, h1, &h1) // h1: main.Human, {Steven 35 1}, 0xc0000044a0
}

func changeName(h Human) {
	h.name = "Daniel"
	h.age = 23
	fmt.Printf("h: %T, %v, %p \n", h, h, &h) // h: main.Human, {Daniel 23 1}, 0xc0000045c0
}
