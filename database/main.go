package main

/*
 * A service for controlling the database.
 */

import (
	"database/sql"
	"fmt"
	"log"

	"time"

	_ "github.com/go-sql-driver/mysql"
	micro "github.com/micro/go-micro"
	"github.com/tjoshum/acca-tracker/database/constants"
	"github.com/tjoshum/acca-tracker/database/handlers"
	"github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
)

var createTablesStatements = []string{
	`CREATE DATABASE IF NOT EXISTS ` + constants.DatabaseName,
	`USE ` + constants.DatabaseName,
	`CREATE TABLE IF NOT EXISTS ` + constants.GameTableName + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		week INT UNSIGNED NOT NULL,
		homeTeam varchar(255),
		awayTeam varchar(255),
		homeScore INT,
		awayScore INT,
		active bool,
		final bool,
		PRIMARY KEY (id)
	)`,
	`CREATE TABLE IF NOT EXISTS ` + constants.UserTableName + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name varchar(255),
		PRIMARY KEY (id)
	)`,
	`CREATE TABLE IF NOT EXISTS ` + constants.BetTableName + ` (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		gameId INT UNSIGNED NOT NULL,
		userId INT UNSIGNED NOT NULL,
		betOn varchar(255),
		spread INT NOT NULL,
		PRIMARY KEY (id)
	)`,
}

func main() {
	// Initialise the database
	db_root, err := sql.Open(constants.DatabaseDriver, constants.ServerString)
	if err != nil {
		log.Fatal("Server open", err)
	}
	defer db_root.Close()
	err = db_root.Ping()
	if err != nil {
		log.Fatal("Ping", err)
	}

	for _, stmt := range createTablesStatements {
		_, err := db_root.Exec(stmt)
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
			"description": "Acca-tracker service to manage access to the database.",
		}),
	)
	service.Init()

	database.RegisterDatabaseServiceHandler(service.Server(), new(handlers.DatabaseHandler))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
