package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"strconv"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
	"golang.org/x/net/context"
)

//type Game database.AddGameRequest // TODO Softer type alias in go 1.9

func sendToDatabase(game database.AddGameRequest) {
	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "no-user",
		"X-From-Id": names.GameDaemon,
	})

	rsp, err := cl.AddGame(ctx, &game)
	if err != nil {
		fmt.Println("AddGame failed", err)
	}
	if rsp.Error != database.ErrorCode_SUCCESS {
		fmt.Println("AddGame failed with response code", rsp.Error.String())
	}
	fmt.Println("Returning from sendToDatabase")
}

type nflRawStruct struct {
	SS [][]string `json:"ss"`
}

type nflGame []string

func (g nflGame) GetAwayTeam() database.TeamCode {
	return names.GetTeamCode(g[4])
}

func (g nflGame) GetHomeTeam() database.TeamCode {
	return names.GetTeamCode(g[6])
}

func (g nflGame) GetHomeScore() int32 {
	i, err := strconv.Atoi(g[5])
	if err != nil {
		return -1
	}
	return int32(i)
}

func (g nflGame) GetAwayScore() int32 {
	i, err := strconv.Atoi(g[7])
	if err != nil {
		return -1
	}
	return int32(i)
}

func (g nflGame) GetWeek() int32 {
	i, err := strconv.Atoi(g[12])
	if err != nil {
		return 0
	}
	return int32(i)
}

func parseJson(jsonStr string) (error, []database.AddGameRequest) {
	str := strings.Replace(jsonStr, ",,", ",\"\",", -1)
	str = strings.Replace(str, ",,", ",\"\",", -1)

	var n nflRawStruct
	err := json.Unmarshal([]byte(str), &n)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return err, nil
	}

	var games []database.AddGameRequest
	for _, strs := range n.SS {
		g := nflGame(strs)
		games = append(games, database.AddGameRequest{
			Week:      g.GetWeek(),
			HomeTeam:  g.GetHomeTeam(),
			AwayTeam:  g.GetAwayTeam(),
			HomeScore: g.GetHomeScore(),
			AwayScore: g.GetAwayScore(),
			Active:    true,
			Final:     true,
		})
	}
	return nil, games
}

func getJson(url string) (error, []database.AddGameRequest) {
	r, err := (&http.Client{Timeout: 10 * time.Second}).Get(url)
	if err != nil {
		fmt.Println("Error getting json from nfl:", err)
		return err, nil
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return err, nil
	}
	return parseJson(string(body))
}

// Fetches the games for this week and sends them to the database.
func fetchCurrentGames() {
	fmt.Println("Running getjson")
	_, games := getJson("http://www.nfl.com/liveupdate/scorestrip/scorestrip.json")
	fmt.Println("Ret getjson")
	for _, g := range games {
		fmt.Println("Running sendToDatabase")
		sendToDatabase(g)
		fmt.Println("Ret sendToDatabase")
	}
}

func main() {
	cmd.Init()
	fetchCurrentGames()
}
