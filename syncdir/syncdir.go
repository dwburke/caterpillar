package syncdir

import (
	"fmt"

	"github.com/otiai10/copy"
)

func Sync() {
	opt := copy.Options{
		PreserveTimes: true,
	}
	err := copy.Copy("/home/dburke/bin", "/home/dburke/bin_copy", opt)
	fmt.Println(err)
}
