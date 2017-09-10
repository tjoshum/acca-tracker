package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"strings"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	database "github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
	"golang.org/x/net/context"
)

func main() {
	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "no-user",
		"X-From-Id": names.GameDaemon,
	})

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username_stripped := strings.TrimSuffix(username, "\n")
	user_req := database.AddUserRequest{
		Username: username_stripped,
	}
	_, err := cl.AddUser(ctx, &user_req)
	if err != nil {
		log.Fatal("Failed to add user", err)
	}

	fmt.Print("Enter week: ")
	week, _ := reader.ReadString('\n')
	week_stripped := strings.TrimSuffix(week, "\n")
	week_int, err := strconv.Atoi(week_stripped)
	if err != nil {
		log.Fatalf("Failed to convert week '%s'", week_stripped)
	}
	rsp, err := cl.GetWeekGames(ctx, &database.GetWeekGamesRequest{
		Week: int32(week_int),
	})
	if err != nil {
		log.Fatal("Failed to get games", err)
	}
	for _, game := range rsp.Games {
		fmt.Println("Enter 1 or 2 (or 0 for no bet)")
		fmt.Println("1.", game.HomeTeam, "vs 2.", game.AwayTeam)
		choice, _ := reader.ReadString('\n')
		choice_stripped := strings.TrimSuffix(choice, "\n")
		choice_int, err := strconv.Atoi(choice_stripped)

		fmt.Println("Choice stripped:", choice_stripped, "Choice int:", choice_int)

		if err != nil {
			log.Fatalf("Failed to convert choice '%s'", choice_stripped)
		}
		if choice_int != 0 {
			fmt.Println("Enter spread (eg +7, -2.5, or 0 for no spread)")
			spread, _ := reader.ReadString('\n')
			spread_stripped := strings.TrimSuffix(spread, "\n")
			spread_int, err := strconv.Atoi(spread_stripped)
			if err != nil {
				log.Fatalf("Failed to convert spread '%s'", spread_stripped)
			}
			var team database.TeamCode
			if choice_int == 1 {
				team = game.HomeTeam
			} else {
				team = game.AwayTeam
			}
			req := database.AddBetRequest{
				GameId:   game.GameId,
				BetOn:    team,
				Spread:   int32(spread_int),
				Username: username_stripped,
			}
			_, err = cl.AddBet(ctx, &req)
			if err != nil {
				log.Fatalf("Failed to add bet", err)
			}
		}

	}
}
