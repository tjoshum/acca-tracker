#!/usr/bin/env bash

set -x

cd $GOPATH/src/github.com/tjoshum/acca-tracker/webd
go get ./...
go run ./main.go --registry_address $REGISTRY_ADDRESS
