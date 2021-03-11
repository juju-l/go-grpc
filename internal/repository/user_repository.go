package repository

import (
	"gitee.com/vipex/go-grpc/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB = nil
var err error = nil

func initCrudRepo() { // 初始化 db
	var dsn = (*utils.GetAppConfigs()).DbConnStr
	db, err = gorm.Open(postgres.New(postgres.Config{ // https://github.com/go-gorm/postgres
		DSN:                  dsn,  //"host=192.168.253.6 user=postgres password=123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
}
func GetCrudRepo() *gorm.DB {
	if err == nil {
		initCrudRepo()
	}
	return db
}

// 新增
func Add() bool {
	return true
}

// 删除
func Delete() bool {
	return true
}

// 更新
func Update() bool {
	return true
}

// 查询
func Select() []interface{} {
	var rst []interface{}
	return rst
}

// 获取分页数据
func GetQueryableByPage() interface{} {
	return new(interface{})
}

// 分页数据封装
func ConvertPages() interface{} {
	return new(interface{})
}
