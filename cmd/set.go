package cmd

import (
	"github.com/spf13/cobra"

	"github.com/dwburke/copyman/db"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set",
	Long:  `set`,
	Run: func(cmd *cobra.Command, args []string) {
		dbh, err := db.Open()
		if err != nil {
			panic(err)
		}
		defer dbh.Close()

		//f := &TestThing{
		//Foo: "jfajfdsjaklfjdsklajfdlksajflkdsajflkdsajflkdsajflkdsjalkjlkjlkjlkjlkjlkjkl",
		//}

		//if err := dbh.SetObj("h", f); err != nil {
		//panic(err)
		//}

	},
}
