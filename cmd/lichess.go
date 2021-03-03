package cmd

import (
	"github.com/foize/go.sgr"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/chrisbarnes2000/makeutility/models"
	"github.com/imroc/req"
	"github.com/spf13/cobra"
)

// PTZone stand for Pacific Standard Time Zone
var PTZone = `T14:00:00.000-08:00`
// GTZone stand for General Time Zone
var GTZone = `2006-01-02T15:04:05.000-07:00`
// dateDisplayFormat stand for General Time Zone
var dateDisplayFormat = `2006-01-02T15:04`
// Grab an API Key From lichess.org/account/oauth/token/create
var apiKey = os.Getenv("LICHESS_TOKEN")

func init() {
	rootCmd.AddCommand(lichessCommand)
	lichessCommand.AddCommand(createCommand)
	lichessCommand.AddCommand(readCommand)
	// lichessCommand.AddCommand(updateCommand)
	// lichessCommand.AddCommand(deleteCommand)
}

var lichessCommand = &cobra.Command{
	Use:   "lichess",
	Short: "Make API Calls to lichess.org",
	Long:  "Interact with lichess.org's API to manage makeschool chess club events",
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

func getTimeStamp(date string) int {
	// https://play.golang.org/p/ouiDtIVjQI
	t, e := time.Parse(GTZone, date)
	if e != nil {
		panic(e)
	}

	return int(t.UTC().UnixNano() / 1000000)
}

// Apply Template Creates An Html File From the Provided Template (.tmpl) & Processed Data
func applyFlyerTemplate(templateName string, data *models.ChessTournamentFlyer) {
	// Create a New HTML File to Apply the Template to
	templatePath := "./tmpl/" + templateName
	savePath := "./docs/" + data.FullName + " " + string(data.StartsAt.Format(dateDisplayFormat)) + ".html"
	newFile, _ := os.Create(savePath)

	// Create a temparay html template file by parsing the tmpl TempALTe at `path`
	t := template.Must(template.New(templateName).ParseGlob(templatePath))

	// err := t.Execute(os.Stdout, data)
	err := t.Execute(newFile, data)
	if err != nil {
		panic(err)
	}
	sgr.Printf("[fg-green bold]Club Flyer Created [fg-yellow]Avalible at: %s[reset]\n", savePath)
}

func createTournamentAndFlyer(date int) {
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
	applyFlyerTemplate("tournament-flyer.tmpl", tournamentInfo)
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
	// applyFlyerTemplate("tournament-flyer.tmpl", tournamentInfo)
}