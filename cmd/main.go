package main

import (
	"google.golang.org/grpc"
	"net"
	pb "gitee.com/vipex/go-grpc/gen/business/user"
	h "gitee.com/vipex/go-grpc/internal/handler"
)

func main() {
	lis, _ := net.Listen("tcp", ":50051")

	s := grpc.NewServer()
	pb.RegisterUserGrpcServer(s, &h.Server{})

	s.Serve(lis)
}
