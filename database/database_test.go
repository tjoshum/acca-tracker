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
