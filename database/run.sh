#!/usr/bin/env bash

set -ex

chown -R mysql:mysql /var/lib/mysql
service mysql start

cd $GOPATH/src/github.com/tjoshum/acca-tracker/database
go get ./...
go run ./main.go --registry_address $REGISTRY_ADDRESS
