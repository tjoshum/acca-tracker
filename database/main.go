package main

import (
	"database/sql"
	"fmt"
	"log"

	"time"

	_ "github.com/go-sql-driver/mysql"
	micro "github.com/micro/go-micro"
	"github.com/tjoshum/acca-tracker/lib/names"
)

/*
A daemon for populating and maintaining a database of games.
*/

const databaseDriver = "mysql"
const serverString = "root:your_password@tcp(127.0.0.1:3306)/"
const databaseName = "accatrackerdb"
const gameTableName = "games"
const betTableName = "bets"

var createTablesStatements = []string{
	`CREATE DATABASE IF NOT EXISTS ` + databaseName,
	`USE ` + databaseName,
	`CREATE TABLE IF NOT EXISTS ` + gameTableName + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		week INT UNSIGNED NOT NULL,
		homeTeam INT UNSIGNED NOT NULL,
		awayTeam INT UNSIGNED NOT NULL,
		PRIMARY KEY (id)
	)`,
	`CREATE TABLE IF NOT EXISTS ` + betTableName + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		gameId INT UNSIGNED NOT NULL,
		PRIMARY KEY (id)
	)`,
}

func main() {
	// Initialise the database
	db_root, err := sql.Open(databaseDriver, serverString)
	if err != nil {
		log.Fatal("Server open", err)
	}
	defer db_root.Close()
	err = db_root.Ping()
	if err != nil {
		log.Fatal("Ping", err)
	}

	_, err = db_root.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName + ";")
	if err != nil {
		log.Fatal("Database create", err)
	}

	for _, stmt := range createTablesStatements {
		_, err := db_root.Exec(stmt + ";")
		if err != nil {
			log.Fatal("Database/table create", err)
		}
	}

	// Initialise and start the service
	service := micro.NewService(
		micro.Name(names.DatabaseSvc),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"description": "Acca-tracker service to manage access to the databse.",
		}),
	)
	service.Init()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
