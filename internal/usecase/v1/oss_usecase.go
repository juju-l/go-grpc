package v1_usecase

import (
	v1_dao "codeup.aliyun.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	v1_interface "codeup.aliyun.com/vipex/go-grpc/internal/domain/v1/v1.interface"
	v1_repoistory "codeup.aliyun.com/vipex/go-grpc/internal/repository/v1"
)

func (use *OssUseCase) Put(req *v1_dao.OssPutReq) (bool, error) {
	return new(v1_repoistory.OssRepository).Put(req)
}

func (use *OssUseCase) Get(req *v1_dao.OssGetReq) (*v1_dao.Oss, error) {
	return new(v1_repoistory.OssRepository).Get(req)
}

type OssUseCase struct {
	ossInterfaceUse v1_interface.OssInterfaceUse
}
