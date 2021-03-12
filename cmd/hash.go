package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/dwburke/caterpillar/hash"
	"github.com/dwburke/caterpillar/util"
)

func init() {
	hashCmd.Flags().String("output", "", "file to save the json to (defaults to <dir.json>)")

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

		dir := util.TrimSuffix(args[0], "/")
		dir, err := filepath.Abs(dir)
		if err != nil {
			return err
		}

		files, err := hash.HashTree(dir)
		if err != nil {
			return err
		}

		save_file := dir + ".json"

		output_file, _ := cmd.Flags().GetString("output")
		if output_file != "" {
			save_file = filepath.Clean(output_file)
		}

		if err = hash.SaveJson(save_file, files); err != nil {
			return err
		}

		return nil
	},
}
