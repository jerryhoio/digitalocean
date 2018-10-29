#!/bin/bash
PROTO=./proto
OUT=./pkg

PACKAGES=`ls $PROTO`
for pkg in $PACKAGES; do
  protoc --proto_path=$GOPATH/src:. --micro_out=$OUT --go_out=$OUT $PROTO/$pkg/*.proto
done
