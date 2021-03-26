package v1_interface

import v1_dao "gitee.com/vipex/go-grpc/internal/domain/v1/v1.dao"

type OssInterfaceRepo interface {
	Put(req *v1_dao.OssPutReq) (bool, error)
	Get(req *v1_dao.OssGetReq) (*v1_dao.Oss, error)
}