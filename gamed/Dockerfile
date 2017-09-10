FROM ubuntu:16.04

RUN apt-get update && apt-get install -y wget git net-tools

# Install modern go.
RUN wget --quiet https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzfgo1.8.1.linux-amd64.tar.gz && \
    ln /usr/local/go/bin/go /usr/bin/go

CMD $GOPATH/src/github.com/tjoshum/acca-tracker/gamed/run.sh
