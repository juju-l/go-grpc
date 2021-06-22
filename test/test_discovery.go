package main

import (
	"context"
	pb "codeup.aliyun.com/vipex/go-grpc/api/vipex.cc/oauth2/proto"
	pri "codeup.aliyun.com/vipex/go-grpc/internal/domain/interface"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"fmt"
)

func main() {
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	service := micro.NewService(
		micro.Name("cc.vipex.client.o2"),
		micro.Registry(etcdRegistry),
	)

	service.Init()

	userClient := pri.NewUserGrpcService("cc.vipex.service.o2", service.Client())

	rsp, err := userClient.Login(context.TODO(), &pb.User{
		User: "admin", Pswd: "654321",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}
