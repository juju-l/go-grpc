package configs

import "gopkg.in/yaml.v2"
import "io/ioutil"

type EtcdAddrCfg struct {
	Addr string `json:"addr"`
}

var etcdAddrCfg = new(EtcdAddrCfg)

func initEtcdAddr() {
	var _addr, err = ioutil.ReadFile("./configs/etcd/_.yml")
	if err == nil {
		if yaml.Unmarshal(_addr, &etcdAddrCfg) != nil {
			// 出错
		}
	}
}

func GetConfig() *EtcdAddrCfg {
	if etcdAddrCfg.Addr == "" {
		initEtcdAddr()
	}
	return etcdAddrCfg
}
