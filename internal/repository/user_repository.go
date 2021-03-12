package repository

import (
	"gitee.com/vipex/go-grpc/internal/domain/dao"
	interfacepri "gitee.com/vipex/go-grpc/internal/domain/interface"
	"gitee.com/vipex/go-grpc/internal/domain/model"
	"gitee.com/vipex/go-grpc/utils"
)

func (s *UserRepository) Login(req *dao.UserReq) (bool, error) {
	var db = (*utils.GetGlobal())["configs"].(*dao.AppConfigs).GetCrudRepo()
	var userInfo model.User; args := "\"user\" = @User AND \"pswd\" = @Pswd" // where 入参(pg 有转义)
		err := db.Where(args, req).First(&userInfo).Error // FirstOrInit
	if err == nil && userInfo.V == true { return true, nil } // 用户登录验证通过 - 产生 token
	return false, nil
}

type UserRepository struct {
	userInterfaceRepo interfacepri.UserInterfaceRepo
}
