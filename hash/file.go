package hash

import (
	"os"
)

type FileData struct {
	Name     string      `json:"name"`
	Hash     string      `json:"hash"`
	FileMode os.FileMode `json:"file_mode"`
}
