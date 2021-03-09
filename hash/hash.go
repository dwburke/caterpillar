package hash

import (
	// #nosec
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

// taken from https://mrwaggel.be/post/generate-md5-hash-of-a-file-in-golang/
// Thanks to Mr. Waggel for a great example for my need

func Md5File(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filepath.Clean(filePath)) // TODO - make sure this doesn' tbreak what already works
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	// #nosec
	defer file.Close()

	//Open a new hash interface to write to
	// #nosec
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}
