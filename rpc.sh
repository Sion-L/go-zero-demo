#!/bin/bash
cd $1 && goctl rpc protoc $1.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero && cd ..