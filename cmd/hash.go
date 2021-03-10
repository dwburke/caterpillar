package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/karrick/godirwalk"
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
	Run: func(cmd *cobra.Command, args []string) {

		var files []*File

		dir := util.TrimSuffix(args[0], "/")
		dir, err := filepath.Abs(dir)
		if err != nil {
			panic(err)
		}

		err = godirwalk.Walk(dir, &godirwalk.Options{
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

				cwd, err := os.Getwd()
				if err != nil {
					return err
				}

				rel, err := filepath.Rel(cwd, osPathname)
				if err != nil {
					return err
				}

				parent := filepath.Base(dir)
				fmt.Printf("%s\n%s/%s\n", osPathname, parent, rel)

				str, err := hash.Md5File(osPathname)
				if err != nil {
					return err
				}

				fmt.Printf("%s %s %s\n",
					de.ModeType(),
					str,
					osPathname)

				f := &File{Name: fmt.Sprintf("%s/%s", parent, rel), Hash: str}
				files = append(files, f)

				return nil
			},
			Unsorted: false, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
		})

		if err != nil {
			panic(err)
		}

		err = util.JsonWrite(dir+".json", files)
		if err != nil {
			panic(err)
		}

	},
}
