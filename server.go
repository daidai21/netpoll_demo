package main

import (
	"context"
	"github.com/cloudwego/netpoll"
	"time"
)

var eventLoop netpoll.EventLoop

func main() {
	network, address := "tcp", ":8080"

	//创建 Listener
	listener, err := netpoll.CreateListener(network, address)
	if err != nil {
		panic("create netpoll listener failed")
	}

	//创建 EventLoop
	eventLoop, _ = netpoll.NewEventLoop(
		handle,
		netpoll.WithOnPrepare(prepare),
		netpoll.WithReadTimeout(time.Second),
	)

	//运行 Server
	eventLoop.Serve(listener)

	//...阻塞
	time.Sleep(time.Duration(10) * time.Second)

	//关闭 Server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	eventLoop.Shutdown(ctx)
}

var _ netpoll.OnPrepare = prepare
var _ netpoll.OnRequest = handle

func prepare(connection netpoll.Connection) context.Context {
	return context.Background()
}

func handle(ctx context.Context, connection netpoll.Connection) error {
	reader := connection.Reader()
	defer reader.Release()
	msg, _ := reader.ReadString(reader.Len())
	println(msg)
	return nil
}
