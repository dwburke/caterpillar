package hash

import (
	"fmt"
	"os"
	"path/filepath"

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

			rel, err := filepath.Rel(dir, osPathname)
			if err != nil {
				return err
			}

			parent := filepath.Base(dir)

			save_filename := fmt.Sprintf("%s/%s", parent, rel)
			f := &FileData{
				Name:     save_filename,
				FileMode: de.ModeType(),
			}
			files[save_filename] = f

			if b, err := de.IsDirOrSymlinkToDir(); b == true && err == nil {
				return nil
			}
			if de.IsDevice() {
				return nil
			}
			if de.ModeType()&os.ModeSocket != 0 {
				return nil
			}
			if de.ModeType()&os.ModeNamedPipe != 0 {
				return nil
			}
			if de.ModeType()&os.ModeCharDevice != 0 {
				return nil
			}

			md5Str := ""
			md5Str, err = Md5File(osPathname)
			if err != nil {
				return err
			}

			f.Hash = md5Str

			fmt.Printf("%s %s/%s\n",
				md5Str, parent, rel)

			return nil
		},
		Unsorted: false,
	})

	if err != nil {
		return "", nil, err
	}

	return dir, files, nil
}
