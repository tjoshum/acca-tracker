package main

import (
	"testing"

	"github.com/tjoshum/acca-tracker/database/proto"
)

func assertEqual(actual interface{}, expected interface{}, t *testing.T) {
	if actual != expected {
		t.Error("Equality check failed", actual, "!=", expected)
	}
}

func TestJsonParse(t *testing.T) {
	pregameStr := `{"ss":[["Thu","20:00:00","Pregame",,"DAL",,"ARI",,,,"57169",,"PRE0","2017"]]}`
	err, games := parseJson(pregameStr)
	if err != nil {
		t.Error("Error fetching preGameStr", err)
	}
	if len(games) != 1 {
		t.Error("Unexpected number of games", len(games))
	}
	assertEqual(games[0].HomeTeam, database.TeamCode_Arizona, t)
	assertEqual(games[0].AwayTeam, database.TeamCode_Dallas, t)
	assertEqual(games[0].HomeScore, int32(0), t)
	assertEqual(games[0].AwayScore, int32(0), t)
}
