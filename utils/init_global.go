package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"gitee.com/vipex/go-grpc/internal/domain/dao" // 
	"sync"
)

var inited bool = false; var globalDef *map[string]interface{}/*存储器*/ = nil; var rwMutex sync.RWMutex // 

func InitGlobal() {
	rwMutex.Lock(); defer rwMutex.Unlock(); if inited { return } // 提供锁

	_globalDef := make(map[string]interface{})
	configs := new(dao.AppConfigs)

	cfg, err := ioutil.ReadFile("./configs/app_configs.yml") 
	if err == nil {
		if yaml.Unmarshal(cfg, &configs) != nil {
			/* 出错处理 */
		}
	}

	_globalDef["configs"] = configs; globalDef = &_globalDef; inited = true // 
}

func GetAppConfigs() *dao.AppConfigs {
	if globalDef == nil { InitGlobal() }; return (*globalDef)["configs"].(*dao.AppConfigs)
}

func GetGlobal() *map[string]interface{} {
	if globalDef == nil {
		InitGlobal()
	}
	return globalDef
}
