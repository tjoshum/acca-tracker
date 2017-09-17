package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"bufio"
	"strings"

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
	fmt.Printf("Password: ")
	password, err := gopass.GetPasswd()
	if err != nil {
		log.Fatal("Error getting password", err)
	}
	return fmt.Sprintf("SKYBETPASSWORD=%s", password)
}

func main() {
	cmd := exec.Command("./bet_getter/skybet/raw_getter/get-raw-bets.sh")
	cmd.Env = append(os.Environ(), getUserEnv(), getPasswordEnv())
	buf, err := cmd.Output()
	if err != nil {
		log.Fatal(string(buf), err)
	}
	raw_html := string(buf)

	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "no-user",
		"X-From-Id": names.GameDaemon,
	})

	re := regexp.MustCompile("(?s)<div class=\"four-six\">.+?</div>")
	slice := re.FindAllStringSubmatch(raw_html, -1)
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
			Username: getPrompt("Username for this bet"),
		}
		cl.AddBet(ctx, &req)

	}
}
