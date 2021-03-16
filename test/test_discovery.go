package main

import (
	"context"
<<<<<<< HEAD
	v1_proto "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/v1/v1.proto"
	v1_interface "gitee.com/vipex/go-grpc/internal/domain/v1/v1.interface"
=======
	pb "gitee.com/vipex/go-grpc/api/vipex.cc/oauth2/proto"
	pri "gitee.com/vipex/go-grpc/internal/domain/interface"
>>>>>>> remotes/origin/single
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

	userClient := v1_interface.NewUserGrpcService("cc.vipex.service.o2", service.Client())

	rsp, err := userClient.Login(context.TODO(), &pb.User{
		User: "admin", Pswd: "654321",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}
