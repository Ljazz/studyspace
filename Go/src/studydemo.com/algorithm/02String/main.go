// package main

// import "fmt"

// const (
// 	Connected = 0
// 	Disconnected = 1
// 	Unknown = 2
// )

// func main() {
// 	a := 20
// 	c := 200
// 	c = a
// 	fmt.Println("赋值操作，把a赋值给c，所以c的值为：", c)
// 	c += a
// 	fmt.Println("相加和赋值运算符，实际为 c = c + a，所以c的值为：", c)
// 	c -= a
// 	fmt.Println("相减和赋值运算符，实际为 c = c - a，所以c的值为：", c)
// 	c *= a
// 	fmt.Println("相乘和赋值运算符，实际为 c = c * a，所以c的值为：", c)
// 	c /= a
// 	fmt.Println("相除和赋值运算符，实际为 c = c / a，所以c的值为：", c)
// 	c <<= 2
// 	fmt.Println("左移和赋值运算符，实际为 c = c << a，所以c的值为：", c)
// 	c >>= 2
// 	fmt.Println("右移和赋值运算符，实际为 c = c >> a，所以c的值为：", c)
// 	c &= 2
// 	fmt.Println("按位与和赋值运算符，实际为 c = c & a，所以c的值为：", c)
// 	c ^= 2
// 	fmt.Println("按位异或和赋值运算符，实际为 c = c ^ a，所以c的值为：", c)
// 	c |= 2
// 	fmt.Println("按位或和赋值运算符，实际为 c = c | a，所以c的值为：", c)
// }

package main

import "fmt"

const (
	a       = iota             // a = 0
	b                          // b = 1, 隐士使用iota关键字，实际等同于 b = iota
	c                          // c = 2, 实际等同于 c = iota
	d, e, f = iota, iota, iota // d = 3, e = 3, f = 3,同一行赋值相同，此处不能写一个iota
	g       = iota             // g = 4
	h       = "h"              // h = "h"，单独赋值，iota依旧递增为5
	i                          // i = "h"，默认使用上面的赋值，iota依旧递增为6
	j       = iota             // j = 7
)

const z = iota // 每个单独定义的const常量中，iota都会重置，此时z=0

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, z)
}
