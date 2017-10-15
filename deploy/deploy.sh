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

go version
if [ $? -ne "0" ]; then
  echo "Go is not installed. Aborting."
  exit 1
fi

if [ -z $GOPATH ]; then
  # GOPATH not set. Assume we're root.
  export GOPATH=/root/go
fi

if [ -z $DOCKER_ADDRESS ]; then
  # Assume we're running on linux, rather than a docker machine.
  export DOCKER_ADDRESS=$(ifconfig docker0 | sed -n -e 's/.*inet addr:\([0-9]*\.*[0-9]*\.[0-9]\.[0-9]\).*/\1/p')
fi

# Set off some long running jobs in parallel.
./teardown.sh& # For dev, cleanup any old containers.
echo "Building containers..."
docker-compose build database&
docker-compose build webd&
docker-compose build gamed&
docker-compose build skybet& # Just to get this into the cache, so that adding bets is quicker.

# Wait for the long running jobs to finish
for job in `jobs -p`; do
    wait $job
done

docker-compose up -d database
sleep 10
docker-compose up -d webd
docker-compose up -d gamed

if [[ $1 == "test" ]]; then
  # Quick and dirty hack. Test within the database container, so I can connect to it in travis.
  docker exec -it deploy_database_1 bash -c "cd /opt/dev/go/src/github.com/tjoshum/acca-tracker; go test -v ./..."
fi
