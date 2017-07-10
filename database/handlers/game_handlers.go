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

func getIDForUsername(db *sql.DB, username string) (int32, error) {
	id_qstr := fmt.Sprintf(
		"SELECT id FROM %s WHERE name = \"%s\"",
		constants.UserTableName,
		username,
	)
	id_rows, err := db.Query(id_qstr)
	if err != nil {
		log.Println("Error: Get user ID", err)
		return 0, err
	}
	var userID int32
	for id_rows.Next() {
		if err := id_rows.Scan(&userID); err != nil {
			log.Fatal(err)
		}
	}
	return userID, nil
}

func stringToTeamCode(str string) database.TeamCode {
	return database.TeamCode(database.TeamCode_value[str])
}

func (s *DatabaseHandler) AddGame(ctx context.Context, req *database.AddGameRequest, rsp *database.AddGameResponse) error {
	db, err := GetDatabase()
	if err != nil {
		log.Println("Error: GetDatabase", err)
		return nil
	}
	defer db.Close()

	str := fmt.Sprintf(
		"INSERT INTO %s (week, homeTeam, awayTeam, homeScore, awayScore, final) VALUES(%d,\"%s\",\"%s\",%d,%d,%d)",
		constants.GameTableName,
		req.Week,
		req.HomeTeam,
		req.AwayTeam,
		req.HomeScore,
		req.AwayScore,
		0) // TODO Take from request - bool to int somehow.
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
	db, err := GetDatabase()
	if err != nil {
		log.Println("Error: GetDatabase", err)
		return nil
	}
	defer db.Close()

	userID, err := getIDForUsername(db, req.GetUsername())

	str := fmt.Sprintf(
		"INSERT INTO %s (gameId, userId, betOn, spread) VALUES(%d,%d,\"%s\",%d)",
		constants.BetTableName,
		req.GetGameId(),
		userID,
		database.TeamCode_name[int32(req.GetBetOn())],
		req.GetSpread())
	log.Println("Attempting SQL '" + str + "'...")
	_, err = db.Exec(str)
	if err != nil {
		log.Println("Error: Bet table entry", err)
		return err
	}
	return nil
}

func (s *DatabaseHandler) AddUser(ctx context.Context, req *database.AddUserRequest, rsp *database.AddUserResponse) error {
	db, err := GetDatabase()
	if err != nil {
		log.Println("Error: GetDatabase", err)
		return nil
	}
	defer db.Close()

	str := fmt.Sprintf(
		"INSERT INTO %s (name) VALUES(\"%s\")",
		constants.UserTableName,
		req.Username)
	log.Println("Attempting SQL '" + str + "'...")
	_, err = db.Exec(str)
	if err != nil {
		log.Println("Error: User table entry", err)
		return err
	}
	return nil
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
			gameId    int32
			week      int32
			homeTeam  string
			awayTeam  string
			homeScore int32
			awayScore int32
			final     bool
		)
		if err := rows.Scan(&gameId, &week, &homeTeam, &awayTeam, &homeScore, &awayScore, &final); err != nil {
			log.Fatal(err)
		}
		println("Game", gameId, "week", week, "home team", homeTeam, "away team", awayTeam)

		rsp.Games = append(rsp.Games, &database.GetWeekGamesResponse_Game{
			gameId,
			stringToTeamCode(homeTeam),
			stringToTeamCode(awayTeam),
			homeScore,
			awayScore,
			final,
		})
	}

	return nil
}

func (s *DatabaseHandler) GetUserBets(ctx context.Context, req *database.GetUserBetsRequest, rsp *database.GetUserBetsResponse) error {
	db, err := GetDatabase()
	if err != nil {
		log.Println("Error: GetDatabase", err)
		return nil
	}
	defer db.Close()

	userID, err := getIDForUsername(db, req.GetUser())
	bet_qstr := fmt.Sprintf(
		"SELECT * FROM %s where userId = %d",
		constants.BetTableName,
		userID,
	)
	bet_rows, err := db.Query(bet_qstr)
	if err != nil {
		log.Println("Error: Get bets", err)
		return err
	}
	for bet_rows.Next() {
		var (
			betID      int32
			gameId     int32
			scanUserId int32
			betOn      string
			spread     int32
		)
		if err := bet_rows.Scan(&betID, &gameId, &scanUserId, &betOn, &spread); err != nil {
			log.Fatal(err)
		}

		rsp.Bets = append(rsp.Bets, &database.GetUserBetsResponse_Bet{
			betID,
			stringToTeamCode(betOn), // TODO
			spread,
		})
	}

	return nil
}

func (s *DatabaseHandler) GetBetsOnGame(ctx context.Context, req *database.GetBetsOnGameRequest, rsp *database.GetBetsOnGameResponse) error {
	db, err := GetDatabase()
	if err != nil {
		log.Println("Error: GetDatabase", err)
		return nil
	}
	defer db.Close()

	bet_qstr := fmt.Sprintf(
		"SELECT * FROM %s where gameId = %d",
		constants.BetTableName,
		req.GetGameId(),
	)
	bet_rows, err := db.Query(bet_qstr)
	if err != nil {
		log.Println("Error: Get bets", err)
		return err
	}
	for bet_rows.Next() {
		var (
			betID      int32
			gameId     int32
			scanUserId int32
			betOn      string
			spread     int32
		)
		if err := bet_rows.Scan(&betID, &gameId, &scanUserId, &betOn, &spread); err != nil {
			log.Fatal(err)
		}

		rsp.Bets = append(rsp.Bets, &database.GetBetsOnGameResponse_Bet{
			betID,
			gameId,
			scanUserId,
			stringToTeamCode(betOn), // TODO
			spread,
		})
	}

	return nil
}

func (s *DatabaseHandler) GetUserList(ctx context.Context, req *database.GetUserListRequest, rsp *database.GetUserListResponse) error {
	db, err := GetDatabase()
	if err != nil {
		log.Println("Error: GetDatabase", err)
		return nil
	}
	defer db.Close()

	str := fmt.Sprintf(
		"SELECT * FROM %s",
		constants.UserTableName,
	)
	rows, err := db.Query(str)
	if err != nil {
		log.Println("Error: Get user list", err)
		return err
	}

	for rows.Next() {
		var userId int32
		var username string
		if err := rows.Scan(&userId, &username); err != nil {
			log.Fatal(err)
		}
		rsp.Users = append(rsp.Users, &database.GetUserListResponse_User{
			userId,
			username,
		})
	}

	return nil
}
