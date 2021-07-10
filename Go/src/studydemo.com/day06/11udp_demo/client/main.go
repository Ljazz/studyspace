package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UDP client

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8000,
	})
	if err != nil {
		fmt.Println("listen UDP failed, err", err)
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	var reply [1024]byte

	// 接收数据
	for {
		fmt.Println("请输入内容：")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("recv reply msg failed, err:", err)
			return
		}
		fmt.Println("收到回复信息:", string(reply[:n]))
	}
}
