package main

import (
	"testing"

	"github.com/tjoshum/acca-tracker/database/handlers"
	"github.com/tjoshum/acca-tracker/database/proto"
	"golang.org/x/net/context"
)

const user1 = "alice"
const user2 = "bob"

func TestGames(t *testing.T) {
	d := new(handlers.DatabaseHandler)

	homeTeam := database.TeamCode_Carolina
	awayTeam := database.TeamCode_Baltimore

	add_req := &database.AddGameRequest{
		Week:      1,
		HomeTeam:  homeTeam,
		AwayTeam:  awayTeam,
		HomeScore: 0,
		AwayScore: 0,
		Final:     false,
	}
	add_rsp := &database.AddGameResponse{}
	d.AddGame(context.TODO(), add_req, add_rsp)

	get_req := &database.GetWeekGamesRequest{
		Week: 1,
	}
	get_rsp := &database.GetWeekGamesResponse{}
	d.GetWeekGames(context.TODO(), get_req, get_rsp)

	for _, game := range get_rsp.GetGames() {
		if game.HomeTeam != homeTeam {
			t.Errorf("Unexpected home team. Got %s expected %s.", game.HomeTeam, homeTeam)
		}
		if game.AwayTeam != awayTeam {
			t.Errorf("Unexpected home team. Got %s expected %s.", game.HomeTeam, homeTeam)
		}
	}

	// Just to populate the database a bit more:
	add_req.HomeTeam = database.TeamCode_GreenBay
	add_req.AwayTeam = database.TeamCode_Buffalo
	err := d.AddGame(context.TODO(), add_req, add_rsp)
	if err != nil {
		t.Error("Get error on second game add", err)
	}
}

func TestUsers(t *testing.T) {
	d := new(handlers.DatabaseHandler)

	add_req := &database.AddUserRequest{
		Username: user1,
	}
	add_rsp := &database.AddUserResponse{}
	d.AddUser(context.TODO(), add_req, add_rsp)

	get_req := &database.GetUserListRequest{}
	get_rsp1 := &database.GetUserListResponse{}
	d.GetUserList(context.TODO(), get_req, get_rsp1)
	if len(get_rsp1.GetUsers()) != 1 {
		t.Errorf("Too many users returned. Got %d expected %d.", len(get_rsp1.GetUsers()), 1)
	}
	for _, user := range get_rsp1.GetUsers() {
		if user.GetUsername() != user1 {
			t.Errorf("Unexpected user. Got %s expected %s.", user1, user)
		}
	}

	add_req2 := &database.AddUserRequest{
		Username: user2,
	}
	d.AddUser(context.TODO(), add_req2, add_rsp)

	get_rsp2 := &database.GetUserListResponse{}
	d.GetUserList(context.TODO(), get_req, get_rsp2)
	if len(get_rsp2.GetUsers()) != 2 {
		t.Errorf("Too many users returned. Got %d expected %d.", len(get_rsp2.GetUsers()), 2)
	}

	found_user1 := false
	found_user2 := false
	for _, user := range get_rsp2.GetUsers() {
		if user.GetUsername() == user1 {
			found_user1 = true
		} else if user.GetUsername() == user2 {
			found_user2 = true
		} else {
			t.Errorf("Found unexpected user: %s", user)
		}
	}
	if !found_user1 {
		t.Errorf("Failed to find user1")
	}
	if !found_user2 {
		t.Errorf("Failed to find user2")
	}
}

// Must run after TestGames and TestUsers, so that we have populated games and users tables.
func TestBets(t *testing.T) {
	d := new(handlers.DatabaseHandler)

	add_req := &database.AddBetRequest{
		GameId:   1,
		BetOn:    database.TeamCode_Carolina,
		Spread:   -2,
		Username: user1,
	}
	add_rsp := &database.AddBetResponse{}
	d.AddBet(context.TODO(), add_req, add_rsp)

	get_req := &database.GetUserBetsRequest{
		Week: 1,
		User: user1,
	}
	get_rsp := &database.GetUserBetsResponse{}

	d.GetUserBets(context.TODO(), get_req, get_rsp)
	for _, bet := range get_rsp.GetBets() {
		if bet.BetOn != database.TeamCode_Carolina {
			t.Errorf("Unexpected team. Got %s expected %s.", bet.BetOn, database.TeamCode_Carolina)
		}
		if bet.Spread != -2 {
			t.Errorf("Unexpected user. Got %d expected %d.", bet.Spread, -2)
		}
	}

	// Just to populate the database a bit more:
	add_req.Username = user2
	add_req.BetOn = database.TeamCode_Baltimore
	err := d.AddBet(context.TODO(), add_req, add_rsp)
	if err != nil {
		t.Error("Get error on second bet add", err)
	}
	add_req.GameId = 2
	add_req.BetOn = database.TeamCode_Buffalo
	err = d.AddBet(context.TODO(), add_req, add_rsp)
	if err != nil {
		t.Error("Get error on third bet add", err)
	}
	add_req.Username = user1
	add_req.BetOn = database.TeamCode_GreenBay
	err = d.AddBet(context.TODO(), add_req, add_rsp)
	if err != nil {
		t.Error("Get error on forth bet add", err)
	}

}
