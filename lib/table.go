package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	database "github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
)

// Type safety, to stop me from getting mixed up.
type Username string
type BetStatus struct {
	BetOnAndSpread string
	GameStatus     string
}
type BetsOnAGame map[Username]BetStatus
type MyTable map[database.Game]BetsOnAGame // games -> {users -> bets}

// Fetch the map of user ID's to usernames
type UserMap map[int32]Username

func getUserMapping(cl database.DatabaseServiceClient, ctx context.Context) UserMap {
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

func getGames(cl database.DatabaseServiceClient, ctx context.Context, weekNum int32) []*database.Game {
	rsp, err := cl.GetWeekGames(ctx, &database.GetWeekGamesRequest{
		Week: weekNum,
	})
	if err != nil {
		log.Fatal("Error getting games for week", weekNum, err)
	}
	fmt.Println("DEBUG webd getGames: size", len(rsp.GetGames()))
	return rsp.GetGames()
}

func getBetsOnGame(cl database.DatabaseServiceClient, ctx context.Context, gameId int32) []*database.GetBetsOnGameResponse_Bet {
	rsp, err := cl.GetBetsOnGame(ctx, &database.GetBetsOnGameRequest{
		GameId: gameId,
	})
	if err != nil {
		log.Fatal("Error getting bets for game", gameId, err)
	}
	return rsp.GetBets()
}

// Determine whether a bet is winning, losing or not started.
func getGameStatus(game database.Game, bet *database.GetBetsOnGameResponse_Bet) string {
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

// External interface. Create a table of the games and bets for a given week.
func CreateTable(week int32) MyTable {
	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "noone1235",
		"X-From-Id": "some-client",
	})

	userMap := getUserMapping(cl, ctx)
	games := getGames(cl, ctx, week)
	fmt.Println("DEBUG webd Game size", len(games))
	display_table := make(MyTable)
	for _, a_game := range games {
		game := *a_game
		if display_table[game] == nil {
			display_table[game] = make(BetsOnAGame)
		}
		fmt.Println("DEBUG webd Going round game loop", game)
		bets_on_this_game := getBetsOnGame(cl, ctx, a_game.GameId)
		for _, bet := range bets_on_this_game {
			user_str := userMap[bet.GetUserId()]
			fmt.Println("DEBUG webd Pushing ", game, "user:", user_str, "bet:", bet.GetBetOn(), bet.GetSpread())
			displayString := fmt.Sprintf("%s (%d)", bet.BetOn, bet.Spread)
			display_table[game][user_str] = BetStatus{displayString, getGameStatus(game, bet)}
		}
	}

	return display_table
}
