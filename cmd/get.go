package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"

	"github.com/dwburke/copyman/db"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get",
	Long:  `get`,
	Run: func(cmd *cobra.Command, args []string) {
		dbh, err := db.Open()
		if err != nil {
			panic(err)
		}
		defer dbh.Close()

		//var f TestThing
		//err = dbh.GetObj("h", &f)
		//fmt.Printf("%#v\n", f)
	},
}
