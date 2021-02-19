package syncdir

import (
	"fmt"

	"github.com/otiai10/copy"
)

func Sync(source string, destination string) error {
	opt := copy.Options{
		PreserveTimes: true,
	}
	err := copy.Copy(source, destination, opt)
	return err
}
