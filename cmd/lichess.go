package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lichessCommand)
}

var lichessCommand = &cobra.Command{
	Use:   "lichess",
	Short: "Make API Calls to lichess.org",
	Long:  "Interact with lichess.org's API to manage makeschool chess club events",
}
