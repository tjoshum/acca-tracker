FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y ruby-dev build-essential zlib1g-dev

RUN gem install highline mechanize

ADD get-raw-bets.rb get-raw-bets.rb

RUN apt-get install -y vim
