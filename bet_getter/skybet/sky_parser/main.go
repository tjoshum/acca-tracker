package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/howeyc/gopass"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	database "github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
	"golang.org/x/net/context"
)

func getPrompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error getting prompt", err)
	}
	return strings.TrimSuffix(text, "\n")
}

func getUserEnv() string {
	username := getPrompt("SkyBet Username")
	return fmt.Sprintf("SKYBETUSER=%s", username)
}

func getPasswordEnv() string {
	fmt.Printf("SkyBet PIN: ")
	password, err := gopass.GetPasswd()
	if err != nil {
		log.Fatal("Error getting password", err)
	}
	return fmt.Sprintf("SKYBETPASSWORD=%s", password)
}

func userInList(user string, list []*database.GetUserListResponse_User) bool {
	for _, user_struct := range list {
		if user_struct.Username == user {
			return true
		}
	}
	return false
}

func getValidUsername(cl database.DatabaseServiceClient, ctx context.Context) string {
	entered := getPrompt("Name for this bet")
	user_list_response, err := cl.GetUserList(ctx, &database.GetUserListRequest{})
	if err != nil {
		log.Fatal("Failed to get user list", err)
	}
	if userInList(entered, user_list_response.Users) {
		return entered
	} else {
		fmt.Println("Not found")
	}
	fmt.Println(user_list_response.Users)
	return entered
}

func getSpread(one_raw_game string) float64 {

	var spread_int float64
	var err error
	if strings.Contains(one_raw_game, "Handicap") {
		var spread_str string
		handicap_re := regexp.MustCompile(`Handicap ((\+|\-)[0-9\.]+)`)
		match := handicap_re.FindStringSubmatch(one_raw_game)
		if len(match) != 0 {
			spread_str = match[1]
		} else {
			handicap_re2 := regexp.MustCompile(`\(((\+|\-)[0-9\.]+)\)`)
			match2 := handicap_re2.FindStringSubmatch(one_raw_game)
			spread_str = match2[1]
		}
		fmt.Println(" spread", spread_str)
		spread_int, err = strconv.ParseFloat(spread_str, 64)
		if err != nil {
			log.Fatal("Failed to convert spread_str", spread_str, err)
		}
	} else {
		fmt.Println(" no spread")
		spread_int = 0
	}
	return spread_int
}

func main() {
	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "no-user",
		"X-From-Id": names.GameDaemon,
	})

	fmt.Println("Building docker container. This may take some time if the cache is cold...")
	cmd := exec.Command("./bet_getter/skybet/raw_getter/get-raw-bets.sh")
	cmd.Env = append(os.Environ(), getUserEnv(), getPasswordEnv())
	buf, err := cmd.Output()
	if err != nil {
		log.Fatal(string(buf), err)
	}
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(buf))
	sel := doc.Find(".individual-bet")
	sel.Each(func(_ int, one_bet *goquery.Selection) {
		fmt.Println("START OF ACCUMULATOR")
		var bets []*database.AddBetRequest
		one_bet.Find(".four-six").Each(func(_ int, acc_line *goquery.Selection) {
			one_raw_game, _ := acc_line.Html()

			bet_on_re := regexp.MustCompile(`(?s)<h3>.+?([A-Z][A-z ]+[a-z]) .+</h3>`)
			bet_on := bet_on_re.FindStringSubmatch(one_raw_game)[1]
			if len(bet_on) == 0 {
				fmt.Println("Failed to parse bet: ", one_raw_game)
				return
			}
			fmt.Print("Bet on: ", bet_on, " with")
			spread_int := getSpread(one_raw_game)

			game_re := regexp.MustCompile("([A-Z][A-z ]+[a-z]) v ([A-Z][A-z ]+[a-z])")
			game_slice := game_re.FindStringSubmatch(one_raw_game)

			// Put validation here for now
			if names.GetTeamCode(game_slice[1]) == database.TeamCode_NotATeam {
				log.Fatal("Failed to translate home ", game_slice[1])
			}
			if names.GetTeamCode(game_slice[2]) == database.TeamCode_NotATeam {
				log.Fatal("Failed to translate away ", game_slice[2])
			}
			if names.GetTeamCode(bet_on) == database.TeamCode_NotATeam {
				log.Fatal("Failed to translate beton ", bet_on)
			}

			get_game_req := database.GetGameRequest{
				HomeTeam: names.GetTeamCode(game_slice[1]),
				AwayTeam: names.GetTeamCode(game_slice[2]),
			}
			rsp, err := cl.GetGame(ctx, &get_game_req)
			if err != nil {
				log.Fatal("Failed to get game", game_slice[1], game_slice[2], err)
			}
			if rsp.GetGame() == nil {
				log.Fatal("Failed to get not nil game", game_slice[1], game_slice[2])
			}

			bets = append(bets, &database.AddBetRequest{
				GameId: rsp.GetGame().GameId,
				BetOn:  names.GetTeamCode(bet_on),
				Spread: int32(spread_int),
			})

		})
		fmt.Println("END OF ACCUMULATOR")
		username_for_this_accumulator := getPrompt("Name for this accumulator")
		for _, b := range bets {
			b.Username = username_for_this_accumulator
			cl.AddBet(ctx, b)
		}

	})

}
