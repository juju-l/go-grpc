package dao

import (
	"crypto/tls"
	"crypto/x509"
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
)

func (appConfigs *AppConfigs)GetCrudRepo() *gorm.DB { // 初始化 db
	var dsn = appConfigs.DbConnStr
	db, _ := gorm.Open(postgres.New(postgres.Config{ // https://github.com/go-gorm/postgres // 配置说明
		DSN: dsn, //"host=192.168.253.6 user=postgres password=123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	return db
}

func (appConfigs *AppConfigs)GetTlsConfig(certFile, keyFile, caFile string) (*tls.Config, error) { // 获取 tlsConfig
	var tlsConfig = &tls.Config{}; tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert; certPool := x509.NewCertPool() // 
	reqCert, err := tls.LoadX509KeyPair(certFile, keyFile); if err != nil { return nil, nil }; tlsConfig.Certificates = []tls.Certificate{ reqCert }
	caCrt, err := ioutil.ReadFile(caFile); if err != nil { return nil, nil }; certPool.AppendCertsFromPEM(caCrt); tlsConfig.RootCAs = certPool
	return tlsConfig, nil
}

func (appConfigs *AppConfigs)GetOssClient() (*aliOss.Client, error) { // 
	return aliOss.New(appConfigs.Oss.Endpoint, appConfigs.Oss.AccessKeyId, appConfigs.Oss.AccessKeySecret) // 获取 aliOss 客户端
}

type AppConfigs struct {
	Etcd      string    `json:"etcd"` //
	DbConnStr string    `json:"dbConnStr" yaml:"dbConnStr"` // 没有 yaml:"dbConnStr" 标签，返回的对象则不能取到值，除非健名全部小写在 yaml 的配置文件中，否则均 xxx
	Oss       OssConfig `json:"oss"` // oss
}

// func (appConfigs *AppConfigs)Get*() err {
// 	return nil
// }
