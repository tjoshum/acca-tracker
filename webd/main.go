package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/prometheus/common/log"
	"github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
)

type displayBets struct {
	DisplayString string
	Class         string
}

// Type safety, to stop me from getting mixed up.
type Username string
type BetsOnAGame map[Username]displayBets
type MyTable map[database.Game]BetsOnAGame // games -> {users -> bets}

type UserMap map[int32]Username

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

func GetGames(cl database.DatabaseServiceClient, ctx context.Context, weekNum int32) []*database.Game {
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

func GetClassString(game database.Game, bet *database.GetBetsOnGameResponse_Bet) string {
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
			displayString := fmt.Sprintf("%s (%d)", bet.BetOn, bet.Spread)
			display_table[game][user_str] = displayBets{displayString, GetClassString(game, bet)}
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
	Game        database.Game
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

func weekViewHandler(w http.ResponseWriter, r *http.Request) {
	week_string := r.URL.Path[len("/week/"):]
	week, err := strconv.Atoi(week_string)
	if err != nil {
		w.Write([]byte("Failed to find week"))
		return
	}
	localTable := createTable(int32(week))

	d := &HeaderData{
		Title: "NFL Betting Results",
		Week:  week_string,
	}
	renderTemplate(w, "head", d)

	var users []Username // De-duplicated list of users who have placed bets this week.
	workspace := make(map[Username]struct{})
	for _, ubmap := range localTable {
		for user, _ := range ubmap {
			workspace[user] = struct{}{}
		}
	}
	for user, _ := range workspace {
		fmt.Println("DEBUG Adding user to ordered list", user)
		users = append(users, user)
	}
	renderTemplate(w, "table_headings", users)

	rownum := 0
	for game, userToBetMap := range localTable {
		fmt.Println("DEBUG LogAGame", game)
		var bets []displayBets

		for _, user := range users {
			bet, present := userToBetMap[user]
			if present {
				fmt.Println("DEBUG LogBet Present", user)
				bets = append(bets, bet)
			} else {
				fmt.Println("DEBUG LogBet NotPresent", user)
				bets = append(bets, displayBets{"", "notstarted"})
			}
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
	http.HandleFunc("/week/", weekViewHandler)
	http.ListenAndServe(":80", nil)
}
