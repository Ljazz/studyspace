package main

import "fmt"

func main() {

	var i1 interface{} = new(Student)
	s := i1.(Student) //不安全，如果断言失败，会直接panic

	fmt.Println(s)

	var i2 interface{} = new(Student)
	s, ok := i2.(Student) //安全，断言失败，也不会panic，只是ok的值为false
	if ok {
		fmt.Println(s)
	}
}

type Student struct {
}
