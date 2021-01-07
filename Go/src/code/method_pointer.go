package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (c *Rectangle) setVal() {
	c.height = 20
}

func main() {
	p := Rectangle{1, 2}
	s := p
	p.setVal()
	fmt.Println(p.height, s.height)
}
