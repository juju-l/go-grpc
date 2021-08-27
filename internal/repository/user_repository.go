package repository

import (
	"codeup.aliyun.com/vipex/go-grpc/internal/domain/dao"
	pri "codeup.aliyun.com/vipex/go-grpc/internal/domain/interface"
	"codeup.aliyun.com/vipex/go-grpc/internal/domain/model"
	"codeup.aliyun.com/vipex/go-grpc/utils"
)

func (s *UserRepository) Login(req *dao.UserReq) (bool, error) {
	db, _ := (*utils.GetGlobal())["configs"].(*dao.AppConfigs).GetCrudRepo()
	var userInfo model.User; args := "\"user\" = @User AND \"pswd\" = @Pswd" // where 入参(pg 有转义)
		err := db.Where(args, req).First(&userInfo).Error // FirstOrInit
	if err == nil && userInfo.V == true { return true, nil } // 用户登录验证通过 - 产生 token
	return false, nil
}

type UserRepository struct {
	userInterfaceRepo pri.UserInterfaceRepo
}
