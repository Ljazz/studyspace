package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp server端

func processConn(conn net.Conn) {
	// 3. 与客户端通信
	var tmp [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed, err:", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		// 发送数据
		fmt.Print("请回复：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		conn.Write([]byte(msg))
	}
}

func main() {
	// 1. 本地端口启动服务
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:8000 failed, err:", err)
		return
	}
	// 2. 等待别人来跟我建立连接
	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go processConn(conn)
	}
}
