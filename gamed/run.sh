#!/usr/bin/env bash

set -x

cd $GOPATH/src/github.com/tjoshum/acca-tracker
go build ./gamed/main.go
./main --registry consul --registry_address $REGISTRY_ADDRESS:8500

