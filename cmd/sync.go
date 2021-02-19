package cmd

//"github.com/dwburke/copyman/syncdir"
//syncdir.Sync()

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "workflow related subcommands",
	Long:  `workflow related subcommands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}
