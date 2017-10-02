#!/usr/bin/env bash

set -ex

chown -R mysql:mysql /var/lib/mysql
service mysql start
go run $GOPATH/src/github.com/tjoshum/acca-tracker/database/main.go --registry_address $REGISTRY_ADDRESS
