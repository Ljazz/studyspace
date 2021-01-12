package main

import (
	"fmt"
)
type Controller struct {
	M int32
}

type Something interface {
	Get()
	Post()
}

func (c *Controller) Get() {
	fmt.Print("GET")
}

func (c *Controller) Post() {
	fmt.Print("POST")
}

type T struct {
	Controller
}

func (t *T) Get() {
	// new(Controller).Get()
	fmt.Print("T")
}
func (t *T) Post() {
	fmt.Print("T")
}
func main() {
	var something Something
	something = new(T)
	var t T
	t.M = 1
	// t.Controller.M = 1
	something.Get()
}