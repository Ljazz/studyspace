package main

import (
	"errors"
	"fmt"
)

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("黑盒11")
	defer fmt.Println("黑盒22")
	defer fmt.Println("黑盒33")
	fmt.Println("end")
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	// deferDemo()
	c, ok := do("*")
	fmt.Println(c(1, 2), ok)
}
