package main

import (
	"flag"
	pri "gitee.com/vipex/go-grpc/internal/domain/interface"
	srv "gitee.com/vipex/go-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	listenAddr := ""
	flag.StringVar(&listenAddr, "addr", ":8080", "app start listen addr...") // 获取参数值
	flag.Parse()

	lis, _ := net.Listen("tcp", listenAddr)

	s := grpc.NewServer()
	pri.RegisterUserGrpcServer(s, &srv.UserServiceGrpc{})
	reflection.Register(s)

	s.Serve(lis) // 
}
