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

type displayBets struct {
	BetOn  database.TeamCode
	Spread int32
	Class  string
}

// Type safety, to stop me from getting mixed up.
type Username string
type UserMap map[int32]Username
type BetsOnAGame map[Username]displayBets
type MyTable map[database.GetWeekGamesResponse_Game]BetsOnAGame // games -> {users -> bets}

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
	fmt.Println("DEBUG webd GetGames: size", len(rsp.GetGames()))
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

func GetClassString(game database.GetWeekGamesResponse_Game, bet *database.GetBetsOnGameResponse_Bet) string {
	if game.GetActive() {
		if bet.GetBetOn() == game.GetHomeTeam() {
			if (game.HomeScore + bet.Spread) > game.AwayScore {
				return "winning"
			} else {
				return "losing"
			}
		} else { // Bet on the away team
			if (game.AwayScore + bet.Spread) > game.HomeScore {
				return "winning"
			} else {
				return "losing"
			}
		}
	} else {
		return "notstarted"
	}
}

func createTable(week int32) MyTable {
	// TODO Share between requests.
	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "noone1235",
		"X-From-Id": "some-client",
	})

	userMap := GetUserMapping(cl, ctx)
	games := GetGames(cl, ctx, week)
	fmt.Println("DEBUG webd Game size", len(games))
	display_table := make(MyTable)
	for _, a_game := range games {
		game := *a_game
		if display_table[game] == nil {
			display_table[game] = make(BetsOnAGame)
		}
		fmt.Println("DEBUG webd Going round game loop", game)
		bets_on_this_game := GetBetsOnGame(cl, ctx, a_game.GameId)
		for _, bet := range bets_on_this_game {
			user_str := userMap[bet.GetUserId()]

			fmt.Println("DEBUG webd Pushing ", game, "user:", user_str, "bet:", bet.GetBetOn(), bet.GetSpread())
			display_table[game][user_str] = displayBets{bet.GetBetOn(), bet.GetSpread(), GetClassString(game, bet)}
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
	Game        database.GetWeekGamesResponse_Game
	RowColour   string
	Predictions []displayBets
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

func getRowColour(rownum int) string {
	if rownum%2 == 0 {
		return "#D3D3D3"
	} else {
		return "#FFFFFF"
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	localTable := createTable(1)

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

	rownum := 0
	for game, userBetStrMap := range localTable {
		fmt.Println("DEBUG LogAGame", game)
		var bets []displayBets
		for _, abet := range userBetStrMap {
			bets = append(bets, abet)
			fmt.Println("LogABet", abet)
		}
		rd := &RowData{
			Game:        game,
			RowColour:   getRowColour(rownum),
			Predictions: bets,
		}
		renderTemplate(w, "row", rd)
		rownum++
	}

	renderTemplate(w, "foot", "")
}

func main() {
	cmd.Init()
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":80", nil)
}
