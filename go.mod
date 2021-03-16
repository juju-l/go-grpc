module gitee.com/vipex/go-grpc

go 1.16

require google.golang.org/grpc v1.27.0 // direct
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
<<<<<<< HEAD
	github.com/golang/protobuf v1.4.0 // direct
	gopkg.in/yaml.v2 v2.2.4
	google.golang.org/protobuf v1.22.0 // direct
=======
	github.com/golang/protobuf v1.4.0
	gopkg.in/yaml.v2 v2.2.4
	gorm.io/driver/postgres v1.0.8
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // 
>>>>>>> remotes/origin/single
)

require github.com/micro/go-micro/v2 v2.9.1 // 核心框架
require gorm.io/gorm v1.21.3 // orm 框架
