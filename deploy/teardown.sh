#!/usr/bin/env bash

IMAGES=$(docker ps -aq)
docker stop $IMAGES
docker rm $IMAGES