package handlers

import (
	"database/sql"
	"fmt"
	"log"

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
		log.Println("Error: Table entry", err)
		return err
	}

	// TODO db.LastInsertId

	log.Println("Successfully added.")
	return nil
}

func (s *DatabaseHandler) AddBet(ctx context.Context, req *database.AddBetRequest, rsp *database.AddBetResponse) error {
	return nil
}

func (s *DatabaseHandler) AddUser(ctx context.Context, req *database.AddUserRequest, rsp *database.AddUserResponse) error {
	return nil
}

func stringToTeamCode(str string) database.TeamCode {
	return database.TeamCode(database.TeamCode_value[str])
}

func (s *DatabaseHandler) GetWeekGames(ctx context.Context, req *database.GetWeekGamesRequest, rsp *database.GetWeekGamesResponse) error {
	db, err := GetDatabase()
	if err != nil {
		log.Println("Error: GetDatabase", err)
		return nil
	}
	defer db.Close()

	str := fmt.Sprintf(
		"SELECT * FROM %s WHERE week = %d",
		constants.GameTableName,
		req.Week,
	)
	rows, err := db.Query(str)
	if err != nil {
		log.Println("Error: Table entry", err)
		return err
	}

	for rows.Next() {
		var (
			gameId   int32
			week     int32
			homeTeam string
			awayTeam string
		)
		if err := rows.Scan(&gameId, &week, &homeTeam, &awayTeam); err != nil {
			log.Fatal(err)
		}
		println("Game", gameId, "week", week, "home team", homeTeam, "away team", awayTeam)

		rsp.Games = append(rsp.Games, &database.GetWeekGamesResponse_Game{
			gameId,
			stringToTeamCode(homeTeam),
			stringToTeamCode(awayTeam),
		})
	}

	return nil
}

func (s *DatabaseHandler) GetGameBets(ctx context.Context, req *database.GetGameBetsRequest, rsp *database.GetGameBetsResponse) error {
	return nil
}
