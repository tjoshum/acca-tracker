package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/prometheus/common/log"
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
	score := g[5]
	if score == "" {
		return 0
	}
	i, err := strconv.Atoi(score)
	if err != nil {
		log.Fatalf("Failed to translate home score '%s'", score)
		return -1
	}
	return int32(i)
}

func (g nflGame) GetAwayScore() int32 {
	score := g[7]
	if score == "" {
		return 0
	}
	i, err := strconv.Atoi(score)
	if err != nil {
		log.Fatalf("Failed to translate away score '%s'", score)
		return -1
	}
	return int32(i)
}

func (g nflGame) GetWeek() int32 {
	re := regexp.MustCompile("^[A-Z]+([0-9]+)$")
	week := re.FindStringSubmatch(g[12])[1]
	i, err := strconv.Atoi(week)
	if err != nil {
		log.Fatalf("Failed to translate week '%s'", g[12])
		return 0
	}
	return int32(i)
}

func parseJson(jsonStr string) (error, []database.AddGameRequest) {
	fmt.Println("Retreived json:", jsonStr)
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
	_, games := getJson("http://www.nfl.com/liveupdate/scorestrip/scorestrip.json")
	for _, g := range games {
		sendToDatabase(g)
	}
}

func main() {
	cmd.Init()
	fetchCurrentGames()
}
