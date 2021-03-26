package v1_repoistory

import (
	"bytes"
	v1_dao "gitee.com/vipex/go-grpc/internal/domain/v1/v1.dao"
	v1_interface "gitee.com/vipex/go-grpc/internal/domain/v1/v1.interface"
	"gitee.com/vipex/go-grpc/utils"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

func (s *UserRepository) Put(req *v1_dao.OssPutReq) (bool, error) {
	var client, err = utils.GetAppConfigs().GetOssClient(); log.Println(err)
		bucket, err := client.Bucket(req.BucketName); log.Println(err)
		objectAcl := oss.ObjectACL(oss.ACLPublicRead)
		if bucket.PutObject(req.ObjectKey, bytes.NewReader(req.ReqFile), objectAcl) != nil {
			return false, err
		}
	return true, nil
}

func (s *UserRepository) Get(req *v1_dao.OssGetReq) (*v1_dao.Oss, error) {
	var client, err = utils.GetAppConfigs().GetOssClient(); log.Println(err)
	var objKeylist = &v1_dao.Oss{}
		bucket, err := client.Bucket(req.BucketName); log.Println(err)
		obj, err := bucket.ListObjects(oss.MaxKeys(7), oss.Marker("")) // 列出文件
		if err != nil { return nil, err } // 异常退出
		for _, val := range obj.Objects {
			objKeylist.Url = append(objKeylist.Url, val.Key)
		}
	return objKeylist, nil
}

type OssRepository struct {
	ossInterfaceRepo v1_interface.OssInterfaceRepo
}
