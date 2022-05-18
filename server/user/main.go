package main

import (
	"douban/dao"
	"douban/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	dao.InitDB()
	InitServer()
}

func InitServer() {
	fmt.Println("服务端已启动")
	// 监听端口
	lis, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterDoubanServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println(1)
}

type server struct {
	proto.UnimplementedDoubanServer
}
