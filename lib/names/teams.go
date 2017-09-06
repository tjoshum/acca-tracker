package names

import "github.com/tjoshum/acca-tracker/database/proto"

func GetTeamCode(team string) database.TeamCode {
	return lookupArr[team]
}

var lookupArr = map[string]database.TeamCode{
	"NotATeam": database.TeamCode_NotATeam,

	"Cleveland": database.TeamCode_Cleveland,
	"CLE":       database.TeamCode_Cleveland,

	"NewOrleans": database.TeamCode_NewOrleans,
	"NO":         database.TeamCode_NewOrleans,

	"NewEngland": database.TeamCode_NewEngland,
	"NE":         database.TeamCode_NewEngland,

	"Detroit": database.TeamCode_Detroit,
	"DET":     database.TeamCode_Detroit,

	"GreenBay": database.TeamCode_GreenBay,
	"GB":       database.TeamCode_GreenBay,

	"Seattle": database.TeamCode_Seattle,
	"SEA":     database.TeamCode_Seattle,

	"Baltimore": database.TeamCode_Baltimore,
	"BAL":       database.TeamCode_Baltimore,

	"Miami": database.TeamCode_Miami,
	"MIA":   database.TeamCode_Miami,

	"Minnesota": database.TeamCode_Minnesota,
	"MIN":       database.TeamCode_Minnesota,

	"Cincinnati": database.TeamCode_Cincinnati,
	"CIN":        database.TeamCode_Cincinnati,

	"Philadelphia": database.TeamCode_Philadelphia,
	"PHI":          database.TeamCode_Philadelphia,

	"Pittsburgh": database.TeamCode_Pittsburgh,
	"PIT":        database.TeamCode_Pittsburgh,

	"Chicago": database.TeamCode_Chicago,
	"CHI":     database.TeamCode_Chicago,

	"Indianapolis": database.TeamCode_Indianapolis,
	"IND":          database.TeamCode_Indianapolis,

	"NYGiants": database.TeamCode_NYGiants,
	"NYG":      database.TeamCode_NYGiants,

	"Jacksonville": database.TeamCode_Jacksonville,
	"JAX":          database.TeamCode_Jacksonville,

	"Los Angeles Rams": database.TeamCode_LARams,
	"LA":               database.TeamCode_LARams,

	"Los Angeles Chargers": database.TeamCode_LAChargers,
	"LAC": database.TeamCode_LAChargers,

	"KansasCity": database.TeamCode_KansasCity,
	"KC":         database.TeamCode_KansasCity,

	"Tennessee": database.TeamCode_Tennessee,
	"TEN":       database.TeamCode_Tennessee,

	"Carolina": database.TeamCode_Carolina,
	"CAR":      database.TeamCode_Carolina,

	"Arizona": database.TeamCode_Arizona,
	"ARI":     database.TeamCode_Arizona,

	"Denver": database.TeamCode_Denver,
	"DEN":    database.TeamCode_Denver,

	"Dallas": database.TeamCode_Dallas,
	"DAL":    database.TeamCode_Dallas,

	"Houston": database.TeamCode_Houston,
	"HOU":     database.TeamCode_Houston,

	"SanFrancisco": database.TeamCode_SanFrancisco,
	"SF":           database.TeamCode_SanFrancisco,

	"Oakland": database.TeamCode_Oakland,
	"OAK":     database.TeamCode_Oakland,

	"NYJets": database.TeamCode_NYJets,
	"NYJ":    database.TeamCode_NYJets,

	"Washington": database.TeamCode_Washington,
	"WAS":        database.TeamCode_Washington,

	"TampaBay": database.TeamCode_TampaBay,
	"TB":       database.TeamCode_TampaBay,

	"Atlanta": database.TeamCode_Atlanta,
	"ATL":     database.TeamCode_Atlanta,

	"Buffalo": database.TeamCode_Buffalo,
	"BUF":     database.TeamCode_Buffalo,
}
