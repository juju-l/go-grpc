package v1_dao

type OssPutReq struct {
	ReqFile    []byte `json:"user"`
	ObjectKey  string `json:"objectKey"`
	BucketName string `json:"bucketName"`
}

type OssGetReq struct {
	BucketName string `json:"bucketName"`
}

type Oss struct {
	Url []string `json:"url"`
}
