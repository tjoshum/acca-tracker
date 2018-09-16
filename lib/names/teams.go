package names

import "github.com/tjoshum/acca-tracker/database/proto"

func GetTeamCode(team string) database.TeamCode {
	return lookupArr[team]
}

var lookupArr = map[string]database.TeamCode{
	"NotATeam": database.TeamCode_NotATeam,

	"Cleveland Browns": database.TeamCode_Cleveland,
	"Cleveland":        database.TeamCode_Cleveland,
	"CLE":              database.TeamCode_Cleveland,

	"New Orleans Saints": database.TeamCode_NewOrleans,
	"New Orleans":        database.TeamCode_NewOrleans,
	"NewOrleans":         database.TeamCode_NewOrleans,
	"NO":                 database.TeamCode_NewOrleans,

	"New England Patriots": database.TeamCode_NewEngland,
	"New England":          database.TeamCode_NewEngland,
	"NewEngland":           database.TeamCode_NewEngland,
	"NE":                   database.TeamCode_NewEngland,

	"Detroit Lions": database.TeamCode_Detroit,
	"Detroit":       database.TeamCode_Detroit,
	"DET":           database.TeamCode_Detroit,

	"Green Bay Packers": database.TeamCode_GreenBay,
	"Green Bay":         database.TeamCode_GreenBay,
	"GreenBay":          database.TeamCode_GreenBay,
	"GB":                database.TeamCode_GreenBay,

	"Seattle Seahawks": database.TeamCode_Seattle,
	"Seattle":          database.TeamCode_Seattle,
	"SEA":              database.TeamCode_Seattle,

	"Baltimore Ravens": database.TeamCode_Baltimore,
	"Baltimore":        database.TeamCode_Baltimore,
	"BAL":              database.TeamCode_Baltimore,

	"Miami Dolphins": database.TeamCode_Miami,
	"Miami":          database.TeamCode_Miami,
	"MIA":            database.TeamCode_Miami,

	"Minnesota Vikings": database.TeamCode_Minnesota,
	"Minnesota":         database.TeamCode_Minnesota,
	"MIN":               database.TeamCode_Minnesota,

	"Cincinnati Bengals": database.TeamCode_Cincinnati,
	"Cincinnati":         database.TeamCode_Cincinnati,
	"CIN":                database.TeamCode_Cincinnati,

	"Philadelphia Eagles": database.TeamCode_Philadelphia,
	"Philadelphia":        database.TeamCode_Philadelphia,
	"PHI":                 database.TeamCode_Philadelphia,

	"Pittsburgh Steelers": database.TeamCode_Pittsburgh,
	"Pittsburgh":          database.TeamCode_Pittsburgh,
	"PIT":                 database.TeamCode_Pittsburgh,

	"Chicago Bears": database.TeamCode_Chicago,
	"Chicago":       database.TeamCode_Chicago,
	"CHI":           database.TeamCode_Chicago,

	"Indianapolis Colts": database.TeamCode_Indianapolis,
	"Indianapolis":       database.TeamCode_Indianapolis,
	"IND":                database.TeamCode_Indianapolis,

	"New York Giants": database.TeamCode_NYGiants,
	"NYGiants":        database.TeamCode_NYGiants,
	"NYG":             database.TeamCode_NYGiants,

	"Jacksonville Jaguars": database.TeamCode_Jacksonville,
	"Jacksonville":         database.TeamCode_Jacksonville,
	"JAX":                  database.TeamCode_Jacksonville,
	"JAC":                  database.TeamCode_Jacksonville,

	"Los Angeles Rams": database.TeamCode_LARams,
	"LA":               database.TeamCode_LARams,

	"Los Angeles Chargers": database.TeamCode_LAChargers,
	"LAC": database.TeamCode_LAChargers,

	"Kansas City Chiefs": database.TeamCode_KansasCity,
	"Kansas City":        database.TeamCode_KansasCity,
	"KansasCity":         database.TeamCode_KansasCity,
	"KC":                 database.TeamCode_KansasCity,

	"Tennessee Titans": database.TeamCode_Tennessee,
	"Tennessee":        database.TeamCode_Tennessee,
	"TEN":              database.TeamCode_Tennessee,

	"Carolina Panthers": database.TeamCode_Carolina,
	"Carolina":          database.TeamCode_Carolina,
	"CAR":               database.TeamCode_Carolina,

	"Arizona Cardinals": database.TeamCode_Arizona,
	"Arizona":           database.TeamCode_Arizona,
	"ARI":               database.TeamCode_Arizona,

	"Denver Broncos": database.TeamCode_Denver,
	"Denver":         database.TeamCode_Denver,
	"DEN":            database.TeamCode_Denver,

	"Dallas Cowboys": database.TeamCode_Dallas,
	"Dallas":         database.TeamCode_Dallas,
	"DAL":            database.TeamCode_Dallas,

	"Houston Texans": database.TeamCode_Houston,
	"Houston":        database.TeamCode_Houston,
	"HOU":            database.TeamCode_Houston,

	"San Francisco 49ers": database.TeamCode_SanFrancisco,
	"San Francisco":       database.TeamCode_SanFrancisco,
	"SanFrancisco":        database.TeamCode_SanFrancisco,
	"SF":                  database.TeamCode_SanFrancisco,

	"Oakland Raiders": database.TeamCode_Oakland,
	"Oakland":         database.TeamCode_Oakland,
	"OAK":             database.TeamCode_Oakland,

	"New York Jets": database.TeamCode_NYJets,
	"NYJets":        database.TeamCode_NYJets,
	"NYJ":           database.TeamCode_NYJets,

	"Washington Redskins": database.TeamCode_Washington,
	"Washington":          database.TeamCode_Washington,
	"WAS":                 database.TeamCode_Washington,

	"Tampa Bay Buccaneers": database.TeamCode_TampaBay,
	"Tampa Bay":            database.TeamCode_TampaBay,
	"TampaBay":             database.TeamCode_TampaBay,
	"TB":                   database.TeamCode_TampaBay,

	"Atlanta Falcons": database.TeamCode_Atlanta,
	"Atlanta":         database.TeamCode_Atlanta,
	"ATL":             database.TeamCode_Atlanta,

	"Buffalo Bills": database.TeamCode_Buffalo,
	"Buffalo":       database.TeamCode_Buffalo,
	"BUF":           database.TeamCode_Buffalo,
}
