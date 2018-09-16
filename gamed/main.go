package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
//    "os"
//    "regexp"
//	"strconv"
	"strings"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
//	"github.com/prometheus/common/log"
	"github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
	"golang.org/x/net/context"
)

//type Game database.UpdateGameRequest // TODO Softer type alias in go 1.9

func sendToDatabase(game database.UpdateGameRequest) {
	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "no-user",
		"X-From-Id": names.GameDaemon,
	})

	rsp, err := cl.UpdateGame(ctx, &game)
	if err != nil {
		fmt.Println("UpdateGame failed", err)
	}
	if rsp.Error != database.ErrorCode_SUCCESS {
		fmt.Println("UpdateGame failed with response code", rsp.Error.String())
	}
}

type nflGame struct {
	Home aTeam `json:"home"`
	Away aTeam `json:"away"`
}

type nflRawStruct map[string]*nflGame

type aTeam struct {
        Score struct {
            Num1 int `json:"1"`
            Num2 int `json:"2"`
            Num3 int `json:"3"`
            Num4 int `json:"4"`
            Num5 int `json:"5"`
            T    int `json:"T"`
        } `json:"score"`
        Abbr string `json:"abbr"`
        To   int    `json:"to"`
}

func (g nflGame) GetHomeTeam() database.TeamCode {
	return names.GetTeamCode(g.Home.Abbr)
}

func (g nflGame) GetAwayTeam() database.TeamCode {
	return names.GetTeamCode(g.Away.Abbr)
}

func (g nflGame) GetActive() bool {
    return true
//	return ((g[2] != "Pregame") && (g[2] != "Final"))
}

func (g nflGame) GetFinal() bool {
    return false
//    return g[2] == "Final"
}

func (g nflGame) GetHomeScore() int32 {
    return int32(g.Home.To)
}

func (g nflGame) GetAwayScore() int32 {
	return int32(g.Away.To)
}

func (g nflGame) GetWeek() int32 {
    return int32(2)
}

func parseJson(jsonStr string) (error, []database.UpdateGameRequest) {
//	fmt.Println("Retreived json:", jsonStr)
	str := strings.Replace(jsonStr, ",,", ",\"\",", -1)
	str = strings.Replace(str, ",,", ",\"\",", -1)

	var n nflRawStruct
	err := json.Unmarshal([]byte(str), &n)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return err, nil
	}

	var games []database.UpdateGameRequest
	for _, g := range n {
		games = append(games, database.UpdateGameRequest{
			Week:      g.GetWeek(),
			HomeTeam:  g.GetHomeTeam(),
            AwayTeam:  g.GetAwayTeam(),
			HomeScore: g.GetHomeScore(),
			AwayScore: g.GetAwayScore(),
			Active:    g.GetActive(),
			Final:     g.GetFinal(),
		})
	}
	return nil, games
}

func getJson(url string) (error, []database.UpdateGameRequest) {
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
func updateCurrentGames() {
	_, games := getJson("http://www.nfl.com/liveupdate/scores/scores.json?season=2018&seasonType=REG&week=2")
	for _, g := range games {
		sendToDatabase(g)
	}
}

func main() {
	cmd.Init()
	for {
		updateCurrentGames()
		time.Sleep(30 * time.Second)
	}
}
