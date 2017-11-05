package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"

	"github.com/micro/go-micro/cmd"
	"github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/table"
)

// Data to be passed in to the header html template
type HeaderData struct {
	Title string
	Week  string
}

// Data to be passed in to the row html template
type RowData struct {
	Game        database.Game
	RowColour   string
	Predictions []table.BetStatus
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

// Boilerplate, for sorting
type sortableGameArray []database.Game

func (s sortableGameArray) Len() int {
	return len(s)
}
func (s sortableGameArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortableGameArray) Less(i, j int) bool {
	// We want to sort chronologically.
	if s[i].Final != s[j].Final {
		return s[i].Final
	}
	if s[i].Active != s[j].Active {
		return s[i].Active
	}
	return s[i].HomeTeam < s[j].HomeTeam
}

func weekViewHandler(w http.ResponseWriter, r *http.Request) {
	week_string := r.URL.Path[len("/week/"):]
	week, err := strconv.Atoi(week_string)
	if err != nil {
		w.Write([]byte("Failed to find week"))
		return
	}
	localTable := table.CreateTable(int32(week))

	d := &HeaderData{
		Title: "NFL Betting Results",
		Week:  week_string,
	}
	renderTemplate(w, "head", d)

	var users []string // De-duplicated list of users who have placed bets this week.
	workspace := make(map[table.Username]struct{})
	for _, ubmap := range localTable {
		for user, _ := range ubmap {
			workspace[user] = struct{}{}
		}
	}
	for user, _ := range workspace {
		fmt.Println("DEBUG Adding user to ordered list", user)
		users = append(users, string(user))
	}
	sort.Strings(users)
	renderTemplate(w, "table_headings", users)

	game_array := make([]database.Game, len(localTable))
	i := 0
	for game, _ := range localTable {
		game_array[i] = game
		i++
	}
	sort.Sort(sortableGameArray(game_array))

	rownum := 0
	for _, game := range game_array {
		userToBetMap := localTable[game]
		fmt.Println("DEBUG LogAGame", game)
		var bets []table.BetStatus

		for _, user := range users {
			bet, present := userToBetMap[table.Username(user)]
			if present {
				fmt.Println("DEBUG LogBet Present", user)
				bets = append(bets, bet)
			} else {
				fmt.Println("DEBUG LogBet NotPresent", user)
				bets = append(bets, table.BetStatus{"", "notstarted"})
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
