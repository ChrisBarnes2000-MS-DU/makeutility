package cmd

import (
	"github.com/foize/go.sgr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version for Makeutility",
	Long:  "All software has versions. This is Makeutility's",
	Run: func(cmd *cobra.Command, args []string) {
		Version()
	},
}

// Version ... Get The current Version
func Version() float32{
	sgr.Println("[fg-green bold]Makeutility's version[reset] is currently [fg-yellow bold]Alpha v1.1[reset]")
	return 1.1
}