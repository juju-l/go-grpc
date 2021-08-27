package main

import (
	"net"

	v1_interface "codeup.aliyun.com/vipex/go-grpc/internal/domain/v1/v1.interface"
	v1_service "codeup.aliyun.com/vipex/go-grpc/internal/service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, _ := net.Listen("tcp", ":8080")
	s := grpc.NewServer()
	v1_interface.RegisterUserGrpcServer(s, &v1_service.UserServiceGrpc{})
	reflection.Register(s) // 反射
	s.Serve(listen) 
}
