package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	database "github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
	"golang.org/x/net/context"
)

func main() {
	buf, _ := ioutil.ReadFile("/tmp/def.html")
	raw := string(buf)

	re := regexp.MustCompile("(?s)<div class=\"four-six\">.+?</div>")
	slice := re.FindAllStringSubmatch(raw, -1)

	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "no-user",
		"X-From-Id": names.GameDaemon,
	})

	for _, game := range slice {
		one_raw_game := game[0]
		bet_on_re := regexp.MustCompile(`(?s)<h3>.+?([A-Z][A-z ]+[a-z]) (\((.[0-9\.]+)\))?.+</h3>`)
		bet_on := bet_on_re.FindStringSubmatch(one_raw_game)[1]
		fmt.Println("Bet on:", bet_on)

		spread := bet_on_re.FindStringSubmatch(one_raw_game)[3]
		fmt.Println("Spread:", spread)

		game_re := regexp.MustCompile("([A-Z][A-z ]+[a-z]) v ([A-Z][A-z ]+[a-z])")
		game_slice := game_re.FindStringSubmatch(one_raw_game)
		fmt.Println("Home Team:", game_slice[1])
		fmt.Println("Away Team:", game_slice[2])

		req := database.AddBetRequest{
			GameId:   1,
			BetOn:    database.TeamCode_Atlanta,
			Spread:   -1,
			Username: "abc",
		}

		cl.AddBet(ctx, &req)

	}
}
