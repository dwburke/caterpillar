package cmd

import (
	"fmt"
	"strings"

	"github.com/karrick/godirwalk"
	"github.com/spf13/cobra"

	"github.com/dwburke/caterpillar/hash"
)

func init() {
	hashCmd.MarkFlagRequired("dir")

	rootCmd.AddCommand(hashCmd)
}

var hashCmd = &cobra.Command{
	Use:     "hash <dir>",
	Short:   "hash",
	Long:    `hash`,
	Args:    cobra.ExactArgs(1),
	Example: "hash foo/",
	Run: func(cmd *cobra.Command, args []string) {

		//if str, err := hash.Md5File("/etc/hosts"); err != nil {
		//panic(err)
		//} else {
		//fmt.Println(str)
		//}

		//dirname := "some/directory/root"
		err := godirwalk.Walk(args[0], &godirwalk.Options{
			Callback: func(osPathname string, de *godirwalk.Dirent) error {
				// Following string operation is not most performant way
				// of doing this, but common enough to warrant a simple
				// example here:
				if strings.Contains(osPathname, ".git") {
					return godirwalk.SkipThis
				}

				if b, err := de.IsDirOrSymlinkToDir(); b == true && err == nil {
					return nil
				}

				str, err := hash.Md5File(osPathname)
				if err != nil {
					return err
				}

				fmt.Printf("%s %s %s\n",
					de.ModeType(),
					str,
					osPathname)

				return nil
			},
			Unsorted: false, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
		})

		if err != nil {
			panic(err)
		}
	},
}
