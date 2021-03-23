FROM reg.vipex.cc:80/library/golang:1.16-alpine3.12.diy

MAINTAINER lj@vipex.cc
# docker build --no-cache -f Dockerfile -t reg.vipex.cc:80/library/golang:1.16-alpine3.12.pb . # build image

ENV PROTOPATH=api/vipex.cc GITHUB=https://gitee.com/vipex/go-grpc.git GITBRANCH=ver_template PACKAGESID=v1_proto VERSIONSID=v1/v1

RUN apk add --no-cache protoc tree && \
	go get -u github.com/golang/protobuf/protoc-gen-go@v1.2.0 && \
	`#go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 &&` \
	go get -u github.com/micro/micro/v2/cmd/protoc-gen-micro@v2.9.3 \
&& mv $GOPATH/bin/protoc-gen-micro $GOPATH/bin/protoc-gen-go-micro \
&& cp -f $GOPATH/bin/protoc-gen-go $GOPATH/bin/protoc-gen-go-grpc

COPY generatePb.sh ~/
RUN chmod 755 ~/generatePb.sh
CMD [ "~/generatePb.sh" ]

# docker run --rm -it -v `pwd`:/stg -e VERSIONSID= reg.vipex.cc:80/library/golang:1.16-alpine3.12.pb ~/generatePb.sh
# ls internal/domain/v1/v1.interface
# tree api/vipex.cc
