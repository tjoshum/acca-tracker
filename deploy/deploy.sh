#!/usr/bin/env bash

set -ex

cd $(dirname "${BASH_SOURCE[0]}")

# Check pre-requisites.
docker --version
if [ $? -ne "0" ]; then
  echo "Docker is not installed. Aborting."
  exit 1
fi

docker-compose --version
if [ $? -ne "0" ]; then
  echo "Docker-compose is not installed. Aborting."
  exit 1
fi

if [ -z $DOCKER_ADDRESS ]; then
  # Assume we're running on linux, rather than a docker machine.
  export DOCKER_ADDRESS=$(ifconfig docker0 | sed -n -e 's/.*inet addr:\([0-9]*\.*[0-9]*\.[0-9]\.[0-9]\).*/\1/p')
fi

# For dev, cleanup any old containers.
./teardown.sh

docker-compose up -d --build database gamed webd

# Quick and dirty hack. Test within the database container, so I can connect to it in travis.
sleep 5
docker exec -it deploy_database_1 bash -c "cd /opt/dev/go/src/github.com/tjoshum/acca-tracker; go test -v ./..."
