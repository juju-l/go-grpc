package main

import (
	"flag"
	pri "gitee.com/vipex/go-grpc/internal/domain/interface"
	srv "gitee.com/vipex/go-grpc/internal/service"
	"gitee.com/vipex/go-grpc/utils"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	// "fmt"
)

var listenAddr string = ""
	flag.StringVar(&listenAddr, "addr", ":8080", "app start listen addr..."); flag.Parse()

func main() {
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
	pri.RegisterUserGrpcHandler(service.Server(), new(srv.UserGrpcHandler))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
