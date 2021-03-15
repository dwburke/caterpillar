package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.caterpillar.yaml)")
}

var rootCmd = &cobra.Command{
	Use:   "caterpillar",
	Short: "caterpillar is a thing",
	Long:  `Love me`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}
