package cmd

import (
	"github.com/spf13/cobra"

	"github.com/yaskoo/conteto/server"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
