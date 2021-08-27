module codeup.aliyun.com/vipex/go-grpc

go 1.16

require google.golang.org/grpc v1.27.0 // direct
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0 // 

require (
	github.com/golang/protobuf v1.4.0
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	gopkg.in/yaml.v2 v2.2.4
	gorm.io/driver/postgres v1.0.8
)

require github.com/micro/go-micro/v2 v2.9.1 // api 核心框架
require gorm.io/gorm v1.21.3 // orm 框架
require github.com/aliyun/aliyun-oss-go-sdk v2.1.6+incompatible 
// require github.com/shima-park/agollo v1.2.10
