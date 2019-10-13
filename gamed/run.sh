#!/usr/bin/env bash

set -x

cd $GOPATH/src/github.com/tjoshum/acca-tracker
go run ./gamed/main.go --registry_address $REGISTRY_ADDRESS
