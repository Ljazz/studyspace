package client

import (
	"github.com/Ljazz/chatserver/protocol"
)

type Client interface {
	Dial(address string) error      // 用于客户端向服务器端发起连接请求，参数是服务器的地址
	Start()                         // 客户端启动，启动后所有客户端的服务都可以使用
	Close()                         // 关闭客户端
	Send(command interface{}) error // 发送消息，注意参数是任意内容
	SetName(name string) error      // 设置用户名
	SendMess(message string) error  // 发送信息，这时候参数是字符串
	InComing() chan protocol.MessCmd
}
