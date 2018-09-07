#!/usr/bin/env bash

set -ex

cd $(dirname "${BASH_SOURCE[0]}")

if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root"
   exit 1
fi

# Misc pre-requisites
apt-get update
apt-get install -y curl vim

# Install docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt-get update
apt-cache policy docker-ce
apt-get install -y docker-ce
#systemctl status docker

# Install docker-compose
curl -o /usr/local/bin/docker-compose -L "https://github.com/docker/compose/releases/download/1.15.0/docker-compose-$(uname -s)-$(uname -m)"
chmod +x /usr/local/bin/docker-compose
docker-compose -v

# Install a modern version of go
wget --quiet https://storage.googleapis.com/golang/go1.10.3.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz
rm go1.10.3.linux-amd64.tar.gz
ln /usr/local/go/bin/go /usr/bin/go

# Get the acca-tracker
go get -d github.com/tjoshum/acca-tracker || true # Ignore the 'no buildable source files' error
echo "SUCCESS: Acca-tracker is available at $HOME/go/src/github.com/tjoshum/acca-tracker"
