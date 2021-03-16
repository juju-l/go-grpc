package v1_dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func (appConfigs *AppConfigs)GetCrudRepo() *gorm.DB { // 初始化 db
	var dsn = appConfigs.DbConnStr
	db, _ := gorm.Open(postgres.New(postgres.Config{ // https://github.com/go-gorm/postgres // 配置说明
		DSN: dsn, //"host=192.168.253.6 user=postgres password=123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	return db
}

type AppConfigs struct {
	Etcd      string    `json:"etcd"` //
	DbConnStr string    `json:"dbConnStr" yaml:"dbConnStr"` // 没有 yaml:"dbConnStr" 则不能取到值，除非健名全小写在 yaml 配置文件中
	Oss       OssConfig `json:"oss"` // oss
}

func (appConfigs *AppConfigs)GetOssClient() error {
	return nil
}
