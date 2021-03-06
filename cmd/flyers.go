package cmd

import (
	"html/template"
	"os"
	"time"

	"github.com/foize/go.sgr"
	"github.com/chrisbarnes2000/makeutility/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(flyerCommand)
	flyerCommand.AddCommand(clubflyerCommand)
	flyerCommand.AddCommand(tournamentflyerCommand)
}

var flyerCommand = &cobra.Command{
	Use:   "flyer",
	Short: "Create an HTML Flyer",
	Long:  "Create an HTML Flyer From The Provided Template",
}

var clubflyerCommand = &cobra.Command{
	Use:   "club",
	Short: "Create an HTML Flyer From The club-flyer Template",
	Long:  "Create an HTML Flyer For The MakeSchool Chess Club, From The club-flyer Template",
	Run: func(cmd *cobra.Command, args []string) {
		// Creating a struct where only the
		// Name, Content & Year are initialised
		info := models.ChessClubFlyer{Year: 2023, Content: "Come Joine The Make School Chess Club"}

		// This will assign default values to
		// non-initialised variables in the struct info
		info.FillDefaults()

		// printing after assigning default variables to struct
		// sgr.Println(info)

		applyClubFlyerTemplate("club-flyer.tmpl", info)
	},
}

var tournamentflyerCommand = &cobra.Command{
	Use:   "tournament",
	Short: "Create an HTML Flyer From The tournament-flyer Template",
	Long:  "Create an HTML Flyer For The MakeSchool Chess Club Tournaments, From The tournament-flyer Template",
	Run: func(cmd *cobra.Command, args []string) {
		// Creating a struct where only the
		// Name, Content & Year are initialised
		info := models.ChessTournamentFlyer{
			LeftImage:  models.Image{
				SRC: "",
				ALT: "",
			},
			RightImage: models.Image{
				SRC: "",
				ALT: "",
			},
			ID:         "",
			FullName:   "MakeSchool Arena",
			Clock: struct {
				Increment int "json:\"increment\""
				Limit     int "json:\"limit\""
			}{},
			Minutes:            0,
			CreatedBy:          "",
			System:             "",
			SecondsToStart:     0,
			SecondsToFinish:    0,
			IsFinished:         false,
			IsRecentlyFinished: false,
			PairingsClosed:     false,
			StartsAt:           time.Time{},
			NbPlayers:          0,
			Perf: struct {
				Icon     string "json:\"icon\""
				Key      string "json:\"key\""
				Name     string "json:\"name\""
				Position int    "json:\"position\""
			}{},
			Schedule: struct {
				Freq  string "json:\"freq\""
				Speed string "json:\"speed\""
			}{},
			Variant: struct {
				Key   string "json:\"key\""
				Name  string "json:\"name\""
				Short string "json:\"short\""
			}{},
			Duels: []struct {
				ID string "json:\"id\""
				P  []struct {
					N string "json:\"n\""
					R int    "json:\"r\""
					K int    "json:\"k\""
				} "json:\"p\""
			}{},
			Standings: struct {
				Page    int "json:\"page\""
				Players []struct {
					Name   string "json:\"name\""
					Rank   int    "json:\"rank\""
					Rating int    "json:\"rating\""
					Score  int    "json:\"score\""
					Sheet  struct {
						Scores []interface{} "json:\"scores\""
						Total  int           "json:\"total\""
						Fire   bool          "json:\"fire\""
					} "json:\"sheet\""
				} "json:\"players\""
			}{},
			Featured: struct {
				ID       string "json:\"id\""
				Fen      string "json:\"fen\""
				Color    string "json:\"color\""
				LastMove string "json:\"lastMove\""
				White    struct {
					Rank   int    "json:\"rank\""
					Name   string "json:\"name\""
					Rating int    "json:\"rating\""
				} "json:\"white\""
				Black struct {
					Rank   int    "json:\"rank\""
					Name   string "json:\"name\""
					Rating int    "json:\"rating\""
				} "json:\"black\""
			}{},
			Podium: []struct {
				Name   string "json:\"name\""
				Rank   int    "json:\"rank\""
				Rating int    "json:\"rating\""
				Score  int    "json:\"score\""
				Sheet  struct {
					Scores []interface{} "json:\"scores\""
					Total  int           "json:\"total\""
					Fire   bool          "json:\"fire\""
				} "json:\"sheet\""
				Nb struct {
					Game   int "json:\"game\""
					Beserk int "json:\"beserk\""
					Win    int "json:\"win\""
				} "json:\"nb\""
				Performance int "json:\"performance\""
			}{},
			Stats: struct {
				Games         int "json:\"games\""
				Moves         int "json:\"moves\""
				WhiteWins     int "json:\"whiteWins\""
				BlackWins     int "json:\"blackWins\""
				Draws         int "json:\"draws\""
				Berserks      int "json:\"berserks\""
				AverageRating int "json:\"averageRating\""
			}{},
		}

		// This will assign default values to
		// non-initialised variables in the struct info
		// info.FillDefaults()

		// printing after assigning default variables to struct
		sgr.Println(info)

		// applyTournamentFlyerTemplate("tournament-flyer.tmpl", &info)
	},
}

// Apply Template Creates An Html File From the Provided Template (.tmpl) & Processed Data
func applyClubFlyerTemplate(templateName string, data models.ChessClubFlyer) {
	// Create a New HTML File to Apply the Template to
	templatePath := "./tmpl/" + templateName
	savePath := "./docs/" + data.Name + ".html"
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

// dateDisplayFormat stand for General Time Zone
var dateDisplayFormat = `2006-01-02T15:04`

// Apply Template Creates An Html File From the Provided Template (.tmpl) & Processed Data
func applyTournamentFlyerTemplate(templateName string, data *models.ChessTournamentFlyer) {
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
	sgr.Printf("[fg-green bold]Tournament Flyer Created [fg-yellow]Avalible at: %s[reset]\n", savePath)
}
