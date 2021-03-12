package utils

import (
	v1_dao "gitee.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var inited bool = false
var rwMutex sync.RWMutex
var globalDef *map[string]interface{} = nil

func InitGlobal() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	if inited {
		return
	}

	_globalDef := make(map[string]interface{})
	configs := new(v1_dao.AppConfigs)

	_configs, err := ioutil.ReadFile("./configs/app_configs.yml")
	if err == nil {
		if yaml.Unmarshal(_configs, &configs) != nil {
			// 出错
		}
	}

	_globalDef["configs"] = configs

	globalDef = &_globalDef

	inited = true
}

func GetAppConfigs() *v1_dao.AppConfigs {
	if globalDef == nil {
		InitGlobal()
	}
	var appConfigs = (*globalDef)["configs"].(*v1_dao.AppConfigs)
	return appConfigs
}

func GetGlobal() *map[string]interface{} {
	if globalDef == nil {
		InitGlobal()
	}
	return globalDef
}
