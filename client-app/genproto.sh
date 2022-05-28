#!/bin/sh

mkdir -p ./genproto

protoc \
    --proto_path=../protobufs \
    --go_out=genproto \
    --go_opt=paths=source_relative \
    --go-grpc_out=genproto \
    --go-grpc_opt=paths=source_relative \
    base.proto
