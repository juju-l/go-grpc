module gitee.com/vipex/go-grpc

go 1.16

require google.golang.org/grpc v1.27.0 // indirect
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.0 // direct
	gopkg.in/yaml.v2 v2.2.4
	google.golang.org/protobuf v1.22.0 // direct
)

require github.com/micro/go-micro/v2 v2.9.1 // 核心框架
require gorm.io/gorm v1.21.3 // orm 框架
