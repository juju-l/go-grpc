package v1_interface

import v1_dao "codeup.aliyun.com/vipex/go-grpc/internal/domain/v1/v1.dao"

type OssInterfaceUse interface {
	Put(req *v1_dao.OssPutReq) (bool, error)
	Get(req *v1_dao.OssGetReq) (*v1_dao.Oss, error)
}
