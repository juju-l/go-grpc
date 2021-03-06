#!/bin/bash

git clone ${GITHUB} -b ${GITBRANCH} .

interfacePath=internal/domain/${VERSIONSID}interface; interfaceSid=v1_interface; if [ "$VERSIONSID" == "" ]; then interfaceSid=interfacepri; fi
rm -f $interfacePath/*.pb.go

pbList=($(ls ${PROTOPATH}))

for pb in ${pbList[@]}; do
	pkgPrefix=$(echo ${GITHUB}|sed 's#https://##g'|sed 's#.git$##g')/${PROTOPATH}/$pb/${VERSIONSID}proto
			rm -f ${PROTOPATH}/$pb/${VERSIONSID}proto/*.pb*go

	protoc -I . --go-grpc_out=plugins=grpc:. --go-grpc_opt=paths=source_relative ${PROTOPATH}/$pb/${VERSIONSID}proto/*.proto
	protoc -I . --go-micro_out=. --plugin=protoc-gen-micro=$GOPATH/bin/protoc-gen-go-micro --go-micro_opt=paths=source_relative ${PROTOPATH}/$pb/${VERSIONSID}proto/*.proto
	cp -f ${PROTOPATH}/$pb/${VERSIONSID}proto/*.pb*go $interfacePath/
			# protoc -I . --go_out=. --go_opt=paths=source_relative ${PROTOPATH}/$pb/${VERSIONSID}proto/*.proto

	sed -i "/package $PACKAGESID/{s/package $PACKAGESID/package $interfaceSid/g;s/$/\n\nimport $PACKAGESID \"${pkgPrefix//\//\\\/}\"/g}" $interfacePath/*.pb.micro.go
	sed -i -e 's/,omitempty//g' -e 's/proto1/pub/g' ${PROTOPATH}/$pb/${VERSIONSID}proto/*.pb.go
done

cd $interfacePath && \
	ls |grep micro|cut -d '.' -f 1|xargs -I {} mv {}.pb.micro.go {}_micro.pb.go && ls |grep -v _|cut -d '.' -f 1|xargs -I {} mv {}.pb.go {}_grpc.pb.go \
&& cd /stg


echo "package $interfaceSid" > $interfacePath/apiResult_grpc.pb.go; echo "package $interfaceSid" > $interfacePath/apiResult_micro.pb.go # 置空，否则影响 31 行

sed -i -e "/package $PACKAGESID/{s/package $PACKAGESID/package $interfaceSid/g;s/ \/\/ import/\n\nimport $PACKAGESID/g}" $interfacePath/*_grpc.pb.go # 替换包名和导入
sed -i '/\/\/ Reference/,/var _ context/{/\/\/ Reference/!{/var _ context/!d}}' $interfacePath/*_grpc.pb.go; sed -i '9,11d;18d' $interfacePath/*_grpc.pb.go # 删除冗余的 pb 定义

sed -i -E -e "s/\*([A-Z])/\*$PACKAGESID.\1\2\3\4\5/g" -e "s/new\((.*)\)/new\($PACKAGESID.\1\)/g" $interfacePath/*.pb.go # 替换对象引用
	sed -i "/HandlerType/s/$PACKAGESID.//g" $interfacePath/*_grpc.pb.go # 恢复特定的 new 对象

# sed -i \ # 更新到 17 行了
# "/$interfaceSid/{s/$/\n\nimport $PACKAGESID \"gitee.com\/vipex\/go-grpc\/api\/vipex.cc\/oauth2\/${VERSIONSID//\//\\\/}proto\"/g}" \
# $interfacePath/*_micro.pb.go

tree ${PROTOPATH}
ls $interfacePath
