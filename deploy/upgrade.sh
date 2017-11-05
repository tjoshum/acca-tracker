#!/usr/bin/env bash

set -x

cd $(dirname "${BASH_SOURCE[0]}")

function bounce_service {
  local service=$1
  docker stop deploy_${service}_1
  docker rm deploy_${service}_1
  docker-compose up -d --build $service
}

git pull
source env.sh
bounce_service gamed&
bounce_service webd&

# Wait for the long running jobs to finish
for job in `jobs -p`; do
    wait $job
done
