#!/usr/bin/env bash

set -ex

cd $(dirname "${BASH_SOURCE[0]}")

docker build -t skybet:latest .
docker run --env SKYBETUSER=$SKYBETUSER --env SKYBETPASSWORD=$SKYBETPASSWORD skybet:latest
