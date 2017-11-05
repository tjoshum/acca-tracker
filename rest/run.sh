#!/usr/bin/env bash

set -x

cd $GOPATH/src/github.com/tjoshum/acca-tracker/rest
go get ./...
go run *.go
