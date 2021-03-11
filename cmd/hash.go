package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dwburke/caterpillar/hash"
	"github.com/dwburke/caterpillar/util"
)

func init() {
	rootCmd.AddCommand(hashCmd)
}

var hashCmd = &cobra.Command{
	Use:                   "hash <dir>",
	Short:                 "hash",
	Long:                  `hash`,
	Args:                  cobra.ExactArgs(1),
	Example:               "hash foo",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		dir, files, err := hash.HashTree(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("Writing file: %s\n", dir+".json")
		err = util.JsonWrite(dir+".json", files)
		if err != nil {
			return err
		}

		return nil
	},
}
