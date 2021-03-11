package utils

import (
	// 导入包
	"gitee.com/vipex/go-grpc/configs/etcd"
	"sync"
)

var inited bool = false
var rwMutex sync.RWMutex
var globalDef *map[string]interface{} = nil

func InitGlobal() {
	rwMutex.Lock()         // 加锁
	defer rwMutex.Unlock() // 解锁状态
	if inited {
		return
	} // 锁状态

	/*
		globalDef = new(map[string]interface{}) // 实例化对象
			报错 -> panic: assignment to entry in nil map
	*/
	var _globalDef = make(map[string]interface{}) // 间接实例化 globalDef
	// mapper.InitMapper() // 初始化 mapper 的注册

	_globalDef["etcd"] = configs.GetConfig()
	// _globalDef["router"] = router.GetRouter()
	/*
			报错 -> import cycle not allowed
		_globalDef["oauth2Srv"] = mid.GetOauth2()
	*/

	globalDef = &_globalDef

	inited = true
}

func GetGlobal() *map[string]interface{} {
	if globalDef == nil {
		InitGlobal()
	}
	return globalDef
}
