package dao

type OssConfig struct {
	AccessKeyId     string `json:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" yaml:"accessKeySecret"`
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
}
