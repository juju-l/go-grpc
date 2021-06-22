package main

import (
	// "flag"
	pri "codeup.aliyun.com/vipex/go-grpc/internal/domain/interface"
	srv "codeup.aliyun.com/vipex/go-grpc/internal/service"
	"codeup.aliyun.com/vipex/go-grpc/utils"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"fmt"
)

func main() {
	// Service Listen Address Param
	// listenAddr := ""
	// flag.StringVar(&listenAddr, "addr", ":8080", "app start listen addr...");flag.Parse() // 获取参数值
	// fmt.Println(listenAddr)

	appConfigs := *utils.GetAppConfigs()
	tlsConfig, err := appConfigs.GetTlsConfig("configs/.srv.crt", "configs/.srv-private.key", "configs/ca.crt")
	if err != nil {
		fmt.Println(err)
	}

	// New Registry
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(appConfigs.Etcd),
		registry.TLSConfig(tlsConfig),
	)

	// New Service
	service := micro.NewService(
		micro.Name("cc.vipex.service.o2"),
		micro.Version("latest"),
		// micro.Address(listenAddr),
		micro.Address(":8080"),
		micro.Registry(etcdRegistry),
	)
	// Initialise service
	service.Init()

	// Register Handler
	pri.RegisterUserGrpcHandler(service.Server(), new(srv.UserGrpcHandler))
	// pri.Register*Handler(service.Server(), new(srv.*Handler)) // 注册其他服务

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
