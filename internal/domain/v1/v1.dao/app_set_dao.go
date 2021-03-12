package v1_dao

type AppConfigs struct {
	Etcd      string    `json:"etcd"`                       //
	DbConnStr string    `json:"dbConnStr" yaml:"dbConnStr"` // 没有 yaml:"dbConnStr" 则不能取到值，除非健名全小写在 yaml 配置文件中
	Oss       OssConfig `json:"oss"`                        // oss
	Es        string    `json:"es"`
	Redis     string    `json:"redis"`
	ApolloCtr string    `json:"apolloCtr" yaml:"apolloCtr"` // apollo 中央配置中心地址信息
	//
	GoMicro bool `json:"goMicro" yaml:"goMicro"` // 是否启用 goMicro 框架，默认启用 后端默认均为 go-gin 框架
	//
	Test string `json:"test"`
}
