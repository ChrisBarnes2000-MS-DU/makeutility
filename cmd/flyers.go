package cmd

import (
	"html/template"
	"os"

	"github.com/foize/go.sgr"
	"github.com/chrisbarnes2000/makeutility/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(flyerCommand)
}

var flyerCommand = &cobra.Command{
	Use:   "flyer",
	Short: "Create an HTML Flyer",
	Long:  "Create an HTML Flyer From The Provided Template",
	Run: func(cmd *cobra.Command, args []string) {
		// Creating a struct where only the
		// Name, Content & Year are initialised
		info := models.ChessClubFlyer{Year: 2023, Content: "Come Joine The Make School Chess Club"}

		// This will assign default values to
		// non-initialised variables in the struct info
		info.FillDefaults()

		// printing after assigning default variables to struct
		// sgr.Println(info)

		applyTemplate("club-flyer.tmpl", info)
	},
}

// Apply Template Creates An Html File From the Provided Template (.tmpl) & Processed Data
func applyTemplate(templateName string, data models.ChessClubFlyer) {
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
