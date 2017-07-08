package main

import (
	"testing"

	"github.com/tjoshum/acca-tracker/database/handlers"
	"github.com/tjoshum/acca-tracker/database/proto"
	"golang.org/x/net/context"
)

func TestGames(t *testing.T) {
	d := new(handlers.DatabaseHandler)

	homeTeam := database.TeamCode_Carolina
	awayTeam := database.TeamCode_Baltimore

	add_req := &database.AddGameRequest{
		Week:     1,
		HomeTeam: homeTeam,
		AwayTeam: awayTeam,
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
}

func TestUsers(t *testing.T) {
	user1 := "alice"
	user2 := "bob"
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
		if user != user1 {
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
		if user == user1 {
			found_user1 = true
		} else if user == user2 {
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
