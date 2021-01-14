package main

import "fmt"

func main(){
	const (
		_ := iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		KB = 1 << (10 * iota)
	)
}