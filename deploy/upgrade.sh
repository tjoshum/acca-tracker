#!/usr/bin/env bash

set -x

cd $(dirname "${BASH_SOURCE[0]}")

git pull

docker stop deploy_gamed_1
docker stop deploy_webd_1
docker rm deploy_gamed_1
docker rm deploy_webd_1

source env.sh
docker-compose up -d --build webd
docker-compose up -d --build gamed
