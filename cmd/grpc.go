package main

import (
	"flag"
	pri "codeup.aliyun.com/vipex/go-grpc/internal/domain/interface"
	srv "codeup.aliyun.com/vipex/go-grpc/internal/service"
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
	pri.RegisterUserGrpcServer(s, &srv.UserServiceGrpc{}) // 注册服务


	reflection.Register(s) // 添加反射的注册，给 grpcui 提供入口

	s.Serve(lis) // 
}
