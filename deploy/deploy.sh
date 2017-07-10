#!/usr/bin/env bash

cd $(dirname "${BASH_SOURCE[0]}")

./teardown.sh

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

docker-compose up -d --build database gamed webd
