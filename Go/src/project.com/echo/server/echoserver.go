package main

import (
	"crypto/rand"
	"crypto/tls"
	"log"
	"net"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("rui.crt", "rui.key")
	if err != nil {
		log.Fatal("server: loadkeys: ", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = time.Now
	config.Rand = rand.Reader

	service := "127.0.0.1:80000"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatal("server: listen: ", err)
	}

	log.Print("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("server: accept: ", err)
			break
		}
		log.Printf("server: accepted from %s", conn.RemoteAddr())

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		log.Print("server: conn: waiting")
		n, err := conn.Read(buf)
		log.Printf("server: conn: wrote %d bytes", n)
		if err != nil {
			log.Printf("server: write: %s", err)
			break
		}
	}
	log.Println("server: conn: closed")
}
