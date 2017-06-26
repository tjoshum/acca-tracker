FROM ubuntu:latest

RUN apt-get update && apt-get install -y wget git net-tools

# Install modern go.
RUN wget https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzfgo1.8.1.linux-amd64.tar.gz && \
    ln /usr/local/go/bin/go /usr/bin/go

# Install and setup MySQL
RUN ["/bin/bash", "-c", "debconf-set-selections <<< 'mysql-server mysql-server/root_password password your_password' && debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password your_password' && apt-get -y install mysql-server"]

# Install go MySQL driver
RUN go get github.com/go-sql-driver/mysql

CMD service mysql start; while true; do sleep 600; done
