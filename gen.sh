#! /bin/bash

# use for protoc to go

cd services/protos
protoc --go_out=../ models.proto
protoc --micro_out=../ --go_out=../ prodservice.proto
protoc-go-inject-tag -input=../models.pb.go
protoc-go-inject-tag -input=../prodservice.pb.go
cd ../../
