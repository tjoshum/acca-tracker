#!/usr/bin/env bash

set -ex

chown -R mysql:mysql /var/lib/mysql
service mysql start

cd $GOPATH/src/github.com/tjoshum/acca-tracker
go run ./database/main.go --registry consul --registry_address $REGISTRY_ADDRESS:8500
