package hash

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/karrick/godirwalk"

	"github.com/dwburke/caterpillar/util"
)

func HashTree(root string) (string, map[string]*FileData, error) {

	files := make(map[string]*FileData)

	dir := util.TrimSuffix(root, "/")
	dir, err := filepath.Abs(dir)
	if err != nil {
		return "", nil, err
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

			str, err := Md5File(osPathname)
			if err != nil {
				return err
			}

			fmt.Printf("%s %s %s\n",
				de.ModeType(),
				str,
				osPathname)

			save_filename := fmt.Sprintf("%s/%s", parent, rel)
			f := &FileData{Name: save_filename, Hash: str}
			files[save_filename] = f

			return nil
		},
		Unsorted: false, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})

	if err != nil {
		return "", nil, err
	}

	return dir, files, nil
}
