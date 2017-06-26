#!/usr/bin/env bash

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

docker-compose up -d --build database