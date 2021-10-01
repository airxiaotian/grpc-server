#!/bin/bash
rm -rf ./interfaces/proto
mkdir -p ./interfaces/proto
protoc -I=./proto/purchase --go_out=plugins=grpc:./interfaces/proto proto/purchase/*.proto
