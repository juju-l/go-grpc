package main

import (
	interfacepri "gitee.com/vipex/go-grpc/internal/domain/interface"
	srv "gitee.com/vipex/go-grpc/internal/service"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("cc.vipex.service.o2"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	interfacepri.RegisterUserGrpcHandler(service.Server(), new(srv.UserGrpcHandler))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
