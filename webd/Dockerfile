FROM ubuntu:latest

RUN apt-get update && apt-get install -y wget git net-tools

# Install modern go.
RUN wget https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.8.1.linux-amd64.tar.gz && \
    ln /usr/local/go/bin/go /usr/bin/go

# For now, hardcode the program path and address
CMD cd $GOPATH/src/github.com/tjoshum/acca-tracker/webd; go run main.go --registry_address $REGISTRY_ADDRESS