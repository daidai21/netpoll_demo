package main

import (
	"github.com/cloudwego/netpoll"
	"time"
)

func main() {
	network, address, timeout := "tcp", "127.0.0.1:8080", 50*time.Millisecond

	//创建 Dialer
	dialer := netpoll.NewDialer()
	conn, err := dialer.DialConnection(network, address, timeout)
	if err != nil {
		panic("dial netpoll connection failed")
	}

	//写入和发送消息
	writer := conn.Writer()
	writer.WriteString("hello world")
	writer.Flush()
}
