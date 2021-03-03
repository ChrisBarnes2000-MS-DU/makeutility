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
		sgr.Println("[fg-green bold]Makeutility's version[reset] is currently [fg-yellow bold]Alpha 1.0[reset]")
	},
}
