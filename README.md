# Acca Tracker

TravisCI Status: [![Build Status](https://travis-ci.org/tjoshum/acca-tracker.svg?branch=master)](https://travis-ci.org/tjoshum/acca-tracker)

## Prerequisites
A freshly install instance of Ubuntu 16.04 LTS, with an internet connection.

## Installation
Copy the installer script from `deploy/install.sh` onto a freshly installed instance of Ubuntu 16.04 LTS, and run it.
You should see output ending:
```
SUCCESS: Acca-tracker is available at /root/go/src/github.com/tjoshum/acca-tracker
```

## Deployment
Change to the directory where acca-tracker was installed, and run the deploy script:
```
./deploy/deploy.sh
```

### Upgrading
To upgrade an already running instance:
```
./deploy/upgrade.sh
```

## The game update daemon
From deployment, the game daemon starts running, updating the database every 30 seconds with the latest results.
You may wish to stop it, with `docker stop deploy_gamed_1` to avoid spamming nfl.com.

## Adding a bet
On a successfully deployed system:

### From SkyBet
```
cd $GOPATH/src/github.com/tjoshum/acca-tracker
go run .bet_getter/skybet/sky_parser/main.go
```

### Manually
```
cd $GOPATH/src/github.com/tjoshum/acca-tracker
go run ./bet_getter/manual/main.go
```

## Viewing the results for a given week
Go to http://<server_address>/week/1 to view week 1's results and bets.
