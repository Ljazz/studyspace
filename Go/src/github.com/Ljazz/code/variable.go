// package main

// var a = "hello"
// var b string = "world"
// var c bool

// func main() {
// 	println(a, b, c)
// 	x := 100
// 	println(&x, x)
// 	x = 200
// 	println(&x, x)
// 	x, y := 300, 400
// 	println(&x, x, y)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// const LENGTH int = 10
// 	// const WIDTH int = 5
// 	// var area int
// 	// const a, b, c = 1, false, "str" // 多重赋值

// 	// area = LENGTH * WIDTH
// 	// fmt.Println("面积为：%d", area)
// 	// println()
// 	// println(a, b, c)

// 	const (
// 		x uint16 = 16
// 		y
// 		s = "abc"
// 		z
// 	)
// 	fmt.Printf("%T, %v", y, y)
// 	fmt.Printf("%T, %v", z, z)
// }

// package main

// import "fmt"

// func main() {
// 	const (
// 		a = iota // 0
// 		b        // 1
// 		c        // 2
// 		d = "ha" // "ha"
// 		e        // "ha"
// 		f = 100  // 100
// 		g        // 100
// 		h = iota // 7
// 		i        // 8
// 	)
// 	fmt.Println(a, b, c, d, e, f, g, h, i)
// }

// // 运行结果：
// // 0 1 2 ha ha 100 100 7 8

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := 100           //int
// 	b := 3.14          //float64
// 	c := true          // bool
// 	d := "Hello World" //string
// 	e := `Ruby`        //string
// 	f := 'A'
// 	fmt.Printf("%T,%b\n", a, a)
// 	fmt.Printf("%T,%f\n", b, b)
// 	fmt.Printf("%T,%t\n", c, c)
// 	fmt.Printf("%T,%s\n", d, d)
// 	fmt.Printf("%T,%s\n", e, e)
// 	fmt.Printf("%T,%d,%c\n", f, f, f)
// 	fmt.Println("-----------------------")
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// 	fmt.Printf("%v\n", d)
// 	fmt.Printf("%v\n", e)
// 	fmt.Printf("%v\n", f)

// }

package main

import (
	"fmt"
)

func main() {
	var x int
	var y float64
	fmt.Println("请输入一个整数,一个浮点数")
	fmt.Scanln(&x, &y) // 读取键盘的输入，通过操作地址，赋值给x和y 阻塞式
	fmt.Printf("x的数值：%d, y的数值：%f\n", x, y)
	fmt.Scanf("%d, %f", &x, &y)
	fmt.Printf("x: %d, y: %f\n", x, y)
}
