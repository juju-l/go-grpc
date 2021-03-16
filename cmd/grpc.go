package main

import (
	"flag"
<<<<<<< HEAD
	pri "gitee.com/vipex/go-grpc/internal/domain/interface"
	srv "gitee.com/vipex/go-grpc/internal/service"
=======
	v1_interface "gitee.com/vipex/go-grpc/internal/domain/v1/v1.interface"
	v1_service "gitee.com/vipex/go-grpc/internal/service/v1"
>>>>>>> remotes/origin/ver_template
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
<<<<<<< HEAD
	pri.RegisterUserGrpcServer(s, &srv.UserServiceGrpc{}) // 注册服务
=======
	v1_interface.RegisterUserGrpcServer(s, &v1_service.UserServiceGrpc{}) // 注册服务
>>>>>>> remotes/origin/ver_template


	reflection.Register(s) // 添加反射的注册，给 grpcui 提供入口

	s.Serve(lis) // 
}
