package utils

import (
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func GetServiceAddr(serviceName string) (string, error) {
	etcdRegistry := etcd.NewRegistry(registry.Addrs(GetAppConfigs().Etcd))
	services, err := etcdRegistry.GetService(serviceName)
	if err != nil {
		// 错误处理
	}
	if node, err := selector.RoundRobin(services)(); err == nil {
		return node.Address, nil
	}
	return serviceName, err // 默认在失败时返回了服务的名称
}
