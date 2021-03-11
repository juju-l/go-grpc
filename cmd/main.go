package main

import (
	"gitee.com/vipex/go-grpc/configs/etcd"
	"gitee.com/vipex/go-grpc/internal/handler"
	"gitee.com/vipex/go-grpc/utils"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	microReg := etcd.NewRegistry(
		registry.Addrs((*utils.GetGlobal())["etcd"].(*configs.EtcdAddrCfg).Addr),
	)

	// New Service
	service := micro.NewService(
		micro.Name("sea2.micro.api.client_manage"),
		micro.Version("latest"),
		micro.Registry(microReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	handler.RegisterUser(service.Server())

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
