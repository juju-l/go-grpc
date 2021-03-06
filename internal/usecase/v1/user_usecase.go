package v1_usecase

import (
	v1_dao "codeup.aliyun.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	v1_interface "codeup.aliyun.com/vipex/go-grpc/internal/domain/v1/v1.interface"
	v1_repoistory "codeup.aliyun.com/vipex/go-grpc/internal/repository/v1"
)

func (use *UserUseCase) Login(req *v1_dao.UserReq) (bool, error) {
	/***
	 * 该 userCase 层主要组装所有的业务逻辑
	 * 该单一业务直接返回，不表示该分层仅执行该内容，复杂业务时，在该服务层中，需继续其他实现
	***/
	return new(v1_repoistory.UserRepository).Login(req)
}

type UserUseCase struct {
	userInterfaceUse v1_interface.UserInterfaceUse
}
