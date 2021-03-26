package v1_service

import (
	"context"
	v1_proto "gitee.com/vipex/go-grpc/api/vipex.cc/aliOss/v1/v1.proto"
	v1_dao "gitee.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	v1_usecase "gitee.com/vipex/go-grpc/internal/usecase/v1"
)

func (srv OssServiceGrpc) Put(ctx context.Context, req *v1_proto.OssPutReq) (*v1_proto.OssPutRst, error) {
	var rst = &v1_proto.OssPutRst{}
	if r, err := new(v1_usecase.OssUseCase).Put(&v1_dao.OssPutReq{ req.ReqFile, req.ObjectKey, req.BucketName }); err == nil {
		rst.Data = r
	}
	return rst, nil
}

func (srv OssServiceGrpc) Get(ctx context.Context, req *v1_proto.OssGetReq) (*v1_proto.OssResult, error) {
	var rst = &v1_proto.OssResult{}
	if r, err := new(v1_usecase.OssUseCase).Get(&v1_dao.OssGetReq{ req.BucketName }); err == nil {
		rst.Data = &v1_proto.Oss{ Url: r.Url }
	}
	return rst, nil
}

type OssServiceGrpc struct {
}
