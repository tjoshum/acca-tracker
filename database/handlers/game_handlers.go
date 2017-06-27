package handlers

import (
	"database/sql"
	"log"

	"fmt"

	"github.com/tjoshum/acca-tracker/database/constants"
	"github.com/tjoshum/acca-tracker/database/proto"
	"golang.org/x/net/context"
)

type DatabaseHandler struct{}

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open(constants.DatabaseDriver, constants.ServerString)
	if err != nil {
		log.Println("Error: Server open", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("Error: Ping", err)
		return nil, err
	}

	_, err = db.Exec("USE " + constants.DatabaseName)
	if err != nil {
		log.Println("Error: Database use", err)
		return nil, err
	}
	return db, nil
}

func (s *DatabaseHandler) AddGame(ctx context.Context, req *database.AddGameRequest, rsp *database.AddGameResponse) error {
	db, err := GetDatabase()
	if err != nil {
		log.Println("Error: GetDatabase", err)
		return nil
	}
	defer db.Close()

	str := fmt.Sprintf(
		"INSERT INTO %s (week, homeTeam, awayTeam) VALUES(%d,\"%s\",\"%s\")",
		constants.GameTableName,
		req.Week,
		req.HomeTeam,
		req.AwayTeam)
	log.Println("Attempting SQL '" + str + "'...")
	_, err = db.Exec(str)
	if err != nil {
		log.Println("Error: Example entry", err)
		return err
	}
	log.Println("Successfully added.")
	return nil
}
