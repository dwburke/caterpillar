package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dwburke/caterpillar/hash"
)

func init() {
	rootCmd.AddCommand(saveCmd)
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save",
	Long:  `save`,
	Run: func(cmd *cobra.Command, args []string) {

		if str, err := hash.Md5File("/etc/hosts"); err != nil {
			panic(err)
		} else {
			fmt.Println(str)
		}

	},
}
