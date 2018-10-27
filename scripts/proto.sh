#!/bin/bash
PROTO=./proto
OUT=./pkg

protoc --proto_path=$GOPATH/src:. --micro_out=$OUT --go_out=$OUT $PROTO/*/*.proto
