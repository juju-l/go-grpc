package v1_repoistory

import (
	v1_dao "gitee.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	v1_interface "gitee.com/vipex/go-grpc/internal/domain/v1/v1.interface"
	v1_model "gitee.com/vipex/go-grpc/internal/domain/v1/v1.model"
	"gitee.com/vipex/go-grpc/utils"
)

func (s *UserRepository) Login(req *v1_dao.UserReq) (bool, error) {
	var db = (*utils.GetGlobal())["configs"].(*v1_dao.AppConfigs).GetCrudRepo()
	var userInfo v1_model.User;args := "\"user\" = @User AND \"pswd\" = @Pswd" // where 入参(pg 有转义)
		err := db.Where(args, login_Req).First(&userInfo).Error // FirstOrInit 第一值
	if err == nil && userInfo.V == true { return true, nil } // 用户登录验证通过 - 产生 token
	return false, nil
}

type UserRepository struct {
	userInterfaceRepo v1_interface.UserInterfaceRepo
}
