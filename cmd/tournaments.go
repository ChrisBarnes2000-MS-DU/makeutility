package cmd

import (
	"github.com/foize/go.sgr"
	"log"
	"os"

	"github.com/chrisbarnes2000/makeutility/models"
	"github.com/imroc/req"
	"github.com/spf13/cobra"
)

// PTZone stand for Pacific Standard Time Zone
var PTZone = `T14:00:00.000-08:00`
// Grab an API Key From lichess.org/account/oauth/token/create
var apiKey = os.Getenv("LICHESS_TOKEN")

func init() {
	lichessCommand.AddCommand(createCommand)
	lichessCommand.AddCommand(readCommand)
	// lichessCommand.AddCommand(updateCommand)
	// lichessCommand.AddCommand(deleteCommand)
}

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "POST Call To lichess.org/api/tournament",
	Long:  "Create A New MakeSchool Chess Club Tournament Through lichess.org/api/tournament",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("You Need To Provied A Date In The Format of: `2021-03-07`")
		}
		// We fist have to convert the provided Date to General Time Then To Pacific Time
		createTournamentAndFlyer(getTimeStamp(args[0] + PTZone))
	},
}

var readCommand = &cobra.Command{
	Use:   "read",
	Short: "GET Call To lichess.org/api/tournament/{id}",
	Long:  "Read A New MakeSchool Chess Club Tournament Through lichess.org/api/tournament",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			panic("You Need To Provied A Valid Tournament ID: zB4D4v7t`")
		}
		getTournamentInfo(args[0])
	},
}

// -------------------------

func createTournamentAndFlyer(date int) {
	apiKey := os.Getenv("LICHESS_TOKEN")

	// Create a new session
	session := req.New()

	header := req.Header{
		"Accept":        "application/json",
		"Authorization": "Bearer " + apiKey,
	}

	params := req.Param{
		"name":           "MakeSchool",
		"clockTime":      10,
		"clockIncrement": 0,
		"minutes":        90,
		"waitMinutes":    60,
		"startDate":      date,
		"variant":        "standard", //"standard" "chess960" "crazyhouse" "antichess" "atomic" "horde" "kingOfTheHill" "racingKings" "threeCheck"
		"rated":          true,
		// "position": ""
		"berserkable": false,
		"streakable":  false,
		"hasChat":     true,
		// "description": "",
		// "password": "",
		// "teambBattleByTeam": "",
		// "conditions.teamMember.teamId": "",
		// "conditions.minRating.rating": 500,
		// "conditions.maxRating.rating": 1200,
		// "conditions.nbRatedGame.nb": 0,
	}

	creationResponse, err := session.Post("https://lichess.org/api/tournament", header, params)
	// Make sure we get a 200 status code from our request
	if creationResponse.Response().StatusCode != 200 || err != nil {
		log.Fatalf("Error Code: %d This request is invalid because we got a non-ok status code...\n %s",
			creationResponse.Response().StatusCode, creationResponse)
	}

	tournamentInfo := &models.ChessTournamentFlyer{}
	creationResponse.ToJSON(tournamentInfo)
	// sgr.Print(tournamentInfo)
	// applyTournamentFlyerTemplate("tournament-flyer.tmpl", tournamentInfo)
}

func getTournamentInfo(tournamentID string) {
	// Create a new session
	session := req.New()

	header := req.Header{
		"Accept":        "application/json",
		"Authorization": "Bearer " + apiKey,
	}

	params := req.Param{
		"page": 1,
	}

	readResponse, err := session.Get("https://lichess.org/api/tournament/"+tournamentID, header, params)
	// Make sure we get a 200 status code from our request
	if readResponse.Response().StatusCode != 200 || err != nil {
		log.Fatalf("Error Code: %d This request is invalid because we got a non-ok status code...\n %s",
			readResponse.Response().StatusCode, readResponse)
	}

	tournamentInfo := &models.ChessTournamentFlyer{}
	readResponse.ToJSON(tournamentInfo)
	sgr.Print(tournamentInfo)
	// applyTournamentFlyerTemplate("tournament-flyer.tmpl", tournamentInfo)
}