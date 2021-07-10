package main

import (
	"github.com/Ljazz/chatserver/server"
)

func main() {
	var s server.Server
	s = server.NewServer()
	s.Listen(":3333")
	s.Start()
}
