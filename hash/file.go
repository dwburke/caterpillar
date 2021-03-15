package hash

import (
	"io/fs"
	"os"
)

type FileData struct {
	Name        string      `json:"name"`
	Hash        string      `json:"hash"`
	FileMode    os.FileMode `json:"file_mode"`
	Permissions fs.FileMode `json:"permissions"`
}
