# Setup the environment variables

if [ -z $GOPATH ]; then
  # GOPATH not set. Assume we're root.
  export GOPATH=/root/go
fi

if [ -z $DOCKER_ADDRESS ]; then
  # Assume we're running on linux, rather than a docker machine.
  export DOCKER_ADDRESS=$(ifconfig docker0 | sed -n -e 's/.*inet addr:\([0-9]*\.*[0-9]*\.[0-9]\.[0-9]\).*/\1/p')
fi
