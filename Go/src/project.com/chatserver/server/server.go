package server

type Server interface {
	Listen(address string) error         // 监听信息的写入
	Broadcast(command interface{}) error // 接收到的信息发送给其它用户
	Start()                              // 启动服务器
	Close()                              // 关闭服务器
}
