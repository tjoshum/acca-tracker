#!/usr/bin/env bash

IMAGES=$(docker ps -aq)
if [[ ! -z $IMAGES ]]; then
  docker stop $IMAGES
  docker rm $IMAGES
fi
