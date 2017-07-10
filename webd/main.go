// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/prometheus/common/log"
	"github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
)

// Type safety, to stop me from getting mixed up.
type Username string
type UserMap map[int32]Username
type BetString string
type GameString string
type BetsOnAGame map[Username]BetString
type MyTable map[GameString]BetsOnAGame //games -> {users -> bets}

func GetUserMapping(cl database.DatabaseServiceClient, ctx context.Context) UserMap {
	rsp, err := cl.GetUserList(ctx, &database.GetUserListRequest{})
	if err != nil {
		log.Fatal("Error getting user list", err)
	}
	ret := make(UserMap)
	for _, user := range rsp.GetUsers() {
		ret[user.GetUserId()] = Username(user.GetUsername())
	}
	return ret
}

func GetGames(cl database.DatabaseServiceClient, ctx context.Context, weekNum int32) []*database.GetWeekGamesResponse_Game {
	rsp, err := cl.GetWeekGames(ctx, &database.GetWeekGamesRequest{
		Week: weekNum,
	})
	if err != nil {
		log.Fatal("Error getting games for week", weekNum, err)
	}
	return rsp.GetGames()
}

func GetBetsOnGame(cl database.DatabaseServiceClient, ctx context.Context, gameId int32) []*database.GetBetsOnGameResponse_Bet {
	rsp, err := cl.GetBetsOnGame(ctx, &database.GetBetsOnGameRequest{
		GameId: gameId,
	})
	if err != nil {
		log.Fatal("Error getting bets for game", gameId, err)
	}
	return rsp.GetBets()
}

func create_strings(week int32) MyTable {
	// TODO Share between requests.
	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "noone1235",
		"X-From-Id": "some-client",
	})

	userMap := GetUserMapping(cl, ctx)
	games := GetGames(cl, ctx, week)

	display_table := make(MyTable)
	for _, a_game := range games {
		bets_on_this_game := GetBetsOnGame(cl, ctx, a_game.GameId)
		for _, bet := range bets_on_this_game {
			game_str := GameString(fmt.Sprintf(
				"%s %d - %d %s",
				a_game.GetAwayTeam(), a_game.GetAwayScore(),
				a_game.GetHomeScore(), a_game.GetHomeTeam()))
			if display_table[game_str] == nil {
				display_table[game_str] = make(BetsOnAGame)
			}
			user_str := userMap[bet.GetUserId()]
			bet_str := BetString(fmt.Sprintf(
				"%s (%d)",
				bet.GetBetOn(),
				bet.GetSpread()))
			display_table[game_str][user_str] = bet_str
		}
	}

	return display_table
}

// Data to be passed in to the header html template
type HeaderData struct {
	Title string
	Week  string
}

// Data to be passed in to the row html template
type RowData struct {
	Headline    GameString
	Predictions []BetString
}

func renderTemplate(w http.ResponseWriter, tmpl string, d interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		fmt.Println("Error renderTemplate ParseFiles: ", err.Error())
	}
	err = t.Execute(w, d)
	if err != nil {
		fmt.Println("Error renderTemplate Execute: ", err.Error())
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	localTable := create_strings(1)

	d := &HeaderData{
		Title: "NFL Betting Results",
		Week:  "1",
	}
	renderTemplate(w, "head", d)

	// TODO De-duplicate
	var users []Username // Users who have placed bets this week.
	for _, ub := range localTable {
		for user := range ub {
			users = append(users, user)
		}
		break
	}
	renderTemplate(w, "table_headings", users)

	for g, ub := range localTable {
		var bets []BetString
		for _, abet := range ub {
			bets = append(bets, abet)
		}
		rd := &RowData{
			Headline:    g,
			Predictions: bets,
		}
		renderTemplate(w, "row", rd)
	}

	renderTemplate(w, "foot", "")
}

func main() {
	cmd.Init()
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":80", nil)
}
