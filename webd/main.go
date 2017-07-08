// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UsernameList []string
type BetString string
type GameString string
type UsersBetsForAGame map[string]BetString
type MyTable map[GameString]UsersBetsForAGame //games -> {users -> bets}

func GetUsers() UsernameList {
	return []string{"tom", "sam"}
}

func GetBetsForWeekUser(week int, user string) UsersBetsForAGame {
	gamebets := make(map[string]BetString)
	if user == "tom" {
		gamebets["CIN 2 v 3 BAL"] = "BAL (+1)"
	} else {
		gamebets["CIN 2 v 3 BAL"] = "BAL (+3)"
	}
	return gamebets
}

func create_strings(week int) MyTable {
	users := GetUsers()

	display_table := make(map[GameString]UsersBetsForAGame)
	for _, username := range users {
		gamebets := GetBetsForWeekUser(1, username)
		for game, bets := range gamebets {
			if display_table[GameString(game)] == nil {
				display_table[GameString(game)] = make(UsersBetsForAGame)
			}
			display_table[GameString(game)][username] = bets
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
	Headline    GameString
	Predictions []BetString
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	localTable := create_strings(1)

	d := &HeaderData{
		Title: "NFL Betting Results",
		Week:  "1",
	}
	renderTemplate(w, "head", d)

	var users []string
	for _, ub := range localTable {
		for user := range ub {
			users = append(users, user)
		}
		break
	}
	renderTemplate(w, "table_headings", users)

	for g, ub := range localTable {
		var bets []BetString
		for _, abet := range ub {
			bets = append(bets, abet)
		}
		rd := &RowData{
			Headline:    g,
			Predictions: bets,
		}
		renderTemplate(w, "row", rd)
	}

	renderTemplate(w, "foot", "")
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":80", nil)
}
