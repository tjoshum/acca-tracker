# Acca Tracker

TravisCI Status: [![Build Status](https://travis-ci.org/tjoshum/acca-tracker.svg?branch=master)](https://travis-ci.org/tjoshum/acca-tracker)

## Prerequisites
- Go
- Docker
- Docker Compose

## Getting Started
```
go get https://github.com/tjoshum/acca-tracker
./acca-tracker/deploy/deploy.sh
```

## Starting the game update daemon
TODO

## Manually adding a bet
After deployment:
```
go run ./bet_getter/manual/main.go
```

## Viewing the results for a given week
Go to http://<server_address>/week/1 to view week 1's results and bets.
