
docker run --rm -it -v `pwd`/go-grpc:/stg -v /etc/apk/repositories:/etc/apk/repositories -w /stg -e GOPROXY=https://goproxy.cn,direct alpine:3.12 sh -c "export PATH=\$PATH:/root/go/bin;sh"

apk add --no-cache git go protoc && git clone https://gitee.com/vipex/go-grpc.git . && git checkout develop && rm -f go.mod go.sum api/vipex.cc/oauth2/proto/*.pb.go internal/domain/interface/*.pb.go
go get -u github.com/golang/protobuf/protoc-gen-go@v1.2.0 && `#go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 && `go get -u github.com/micro/micro/v2/cmd/protoc-gen-micro@v2.9.3
mv ~/go/bin/protoc-gen-micro ~/go/bin/protoc-gen-go-micro&&cp -f ~/go/bin/protoc-gen-go ~/go/bin/protoc-gen-go-grpc # 老版本-兼容,新版本不需要执行该行

### protoc -I api/vipex.cc/oauth2/proto --go_out=api/vipex.cc/oauth2/proto --go_opt=paths=source_relative api/vipex.cc/oauth2/proto/*.proto
protoc -I api/vipex.cc/oauth2/proto --go-grpc_out=plugins=grpc:api/vipex.cc/oauth2/proto --go-grpc_opt=paths=source_relative api/vipex.cc/oauth2/proto/*.proto # 开放客户端的注册必须有该部分内容，不能只有 pb 的定义

protoc -I api/vipex.cc/oauth2/proto --go-micro_out=internal/domain/interface --plugin=protoc-gen-micro=~/go/bin/protoc-gen-go-micro --go-micro_opt=paths=source_relative api/vipex.cc/oauth2/proto/*.proto
# protoc -I api/vipex.cc/oauth2/proto --go-grpc_out=internal/domain/interface --plugin=grpc --go-grpc_opt=paths=source_relative api/vipex.cc/oauth2/proto/*.proto # 新版 ok
protoc -I api/vipex.cc/oauth2/proto --go-grpc_out=plugins=grpc:internal/domain/interface --go-grpc_opt=paths=source_relative api/vipex.cc/oauth2/proto/*.proto # 老版本-兼容执行该行

p=`pwd`&&cd internal/domain/interface&&ls |grep micro|cut -d '.' -f 1|xargs -I {} mv {}.pb.micro.go {}_micro.pb.go&&cd $p # 处理 micro 文件的命名
p=`pwd`&&cd internal/domain/interface&&ls|grep -v _|cut -d '.' -f 1|xargs -I {} mv {}.pb.go {}_grpc.pb.go&&cd $p # 老版本-兼容，需特殊执行该部分内容，才能在项目中使用
sed -i 's/,omitempty//g' api/vipex.cc/oauth2/proto/*.pb.go # pb 文件的置零

sed -i -e '/package pb/{s/package pb/package interfacepri/g;s/ \/\/ import/\n\nimport pb/g}' internal/domain/interface/*.pb.go # 老版本-兼容，需特殊执行该部分内容，才能在项目中使用
sed -i '/\/\/ Reference/,/var _ context/{/\/\/ Reference/!{/var _ context/!d}}' internal/domain/interface/*_grpc.pb.go # 老版本-兼容，需特殊执行该部分内容，才能在项目中使用
sed -i '9,10d;18d' internal/domain/interface/*_grpc.pb.go # 老版本-兼容，需特殊执行该部分内容，才能在项目中使用

sed -i -E -e 's/\*([A-Z])/\*pb.\1\2\3\4\5/g' -e 's/new\((.*)\)/new\(pb.\1\)/g' internal/domain/interface/*.pb.go; sed -i '/HandlerType/s/pb.//g' internal/domain/interface/*_grpc.pb.go # 老版本-兼容
sed -i "/interfacepri/{s/$/\n\nimport pb \"gitee.com\/vipex\/go-grpc\/api\/vipex.cc\/oauth2\/proto\"/g}" internal/domain/interface/*_micro.pb.go
