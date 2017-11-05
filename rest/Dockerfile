FROM ubuntu:16.04

RUN apt-get update && apt-get install -y wget git net-tools

# Install modern go.
RUN wget --quiet https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz && \
    ln /usr/local/go/bin/go /usr/bin/go

CMD $GOPATH/src/github.com/tjoshum/acca-tracker/rest/run.sh
