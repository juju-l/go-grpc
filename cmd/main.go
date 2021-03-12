package main

import (
	"flag"
	v1_pri "gitee.com/vipex/go-grpc/internal/domain/v1/v1.interface"
	v1_srv "gitee.com/vipex/go-grpc/internal/service/v1"
	"gitee.com/vipex/go-grpc/utils"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"fmt"
)

func main() {
	// Service Listen Address Param
	listenAddr := ""
	flag.StringVar(&listenAddr, "addr", ":8080", "app start listen addr...");flag.Parse() // 获取参数值
	fmt.Println(listenAddr)

	// New Registry
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs((*utils.GetAppConfigs()).Etcd),
	)

	// New Service
	service := micro.NewService(
		micro.Name("cc.vipex.service.o2"),
		micro.Version("latest"),
		micro.Address(listenAddr),
		micro.Registry(etcdRegistry),
	)
	// Initialise service
	service.Init()

	// Register Handler
	v1_pri.RegisterUserGrpcHandler(service.Server(), new(v1_srv.UserGrpcHandler))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
