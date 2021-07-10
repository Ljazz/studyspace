package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	// 1. 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:8000 failed, err:", err)
		return
	}
	// 2. 发送数据
	reader := bufio.NewReader(os.Stdin)
	var tmp [1024]byte
	for {
		fmt.Print("请说话：")
		msg, _ := reader.ReadString('\n') // 读到换行
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		// 接收数据
		conn.Write([]byte(msg))
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
	conn.Close()
}
