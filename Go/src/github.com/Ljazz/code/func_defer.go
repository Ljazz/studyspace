package main

import (
	"fmt"
)

// func main() {
// 	a := 1
// 	b := 2
// 	defer fmt.Println("b = ", b)
// 	fmt.Println("a = ", a)
// }

// func finished() {
// 	fmt.Println("Finished finding largest")
// }

// func largest(nums []int) {
// 	defer finished()
// 	fmt.Println("Started finding largest")
// 	max := nums[0]
// 	for _, v := range nums {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	fmt.Println("Largest number in", nums, "is", max)
// }

// func main() {
// 	nums := []int{78, 109, 2, 563, 300}
// 	largest(nums)
// }

// 延迟方法

// type person struct {
// 	firstName string
// 	lastName  string
// }

// func (p person) fullName() {
// 	fmt.Printf("%s %s", p.firstName, p.lastName)
// }

// func main() {
// 	p := person{
// 		firstName: "John",
// 		lastName:  "Smith",
// 	}
// 	defer p.fullName()
// 	fmt.Printf("Welcome ")
// }

// 延迟参数

// func printA(a int){
// 	fmt.Println("value of a in deferred function", a)
// }

// func main(){
// 	a := 5
// 	defer printA(a)
// 	a = 10
// 	fmt.Println("value of a in deferred function call", a)
// }

// 堆栈延迟

func main() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
