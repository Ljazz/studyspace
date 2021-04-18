package main

import (
	"fmt"
)

func main() {
	s := "Hello world"
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}
