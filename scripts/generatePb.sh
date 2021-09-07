#!/bin/bash

find ./api/ -name *.go -exec rm -f {} \;

protoc -I . \
--go_out . --go_opt=paths=source_relative \
--go-grpc_out . --go-grpc_opt=paths=source_relative \
--go-grpc_opt=require_unimplemented_servers=false \
api/vipex.cc/*/v1/v1.proto/*.proto


tree api/vipex.cc
#...