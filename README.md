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

## The game update daemon
From deployment, the game daemon starts running, updating the database every 30 seconds with the latest results.
You may wish to stop it, with `docker stop deploy_gamed_1` to avoid spamming nfl.com.

## Manually adding a bet
After deployment:
```
go run ./bet_getter/manual/main.go
```

## Viewing the results for a given week
Go to http://<server_address>/week/1 to view week 1's results and bets.
