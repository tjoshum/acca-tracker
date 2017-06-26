package models

type Game struct {
	Week      int
	HomeTeam  string
	AwayTeam  string
	Final     bool
	HomeScore int
	AwayScore int
}

type Bet struct {
	MyGame *Game
	Team   string
	Spread int
}
