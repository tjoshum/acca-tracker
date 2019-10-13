FROM ubuntu:16.04

RUN apt-get update && apt-get install -y wget git net-tools

# Install modern go.
RUN wget --quiet https://storage.googleapis.com/golang/go1.13.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzfgo1.13.1.linux-amd64.tar.gz && \
    ln /usr/local/go/bin/go /usr/bin/go

# Install and setup MySQL
RUN ["/bin/bash", "-c", "debconf-set-selections <<< 'mysql-server mysql-server/root_password password your_password' && \
     debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password your_password' && \
     apt-get -y install mysql-server"]

# Install go MySQL driver
RUN go get github.com/go-sql-driver/mysql
RUN go get google.golang.org/grpc/status

CMD $GOPATH/src/github.com/tjoshum/acca-tracker/database/run.sh
