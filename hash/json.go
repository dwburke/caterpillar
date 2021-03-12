package hash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dwburke/caterpillar/util"
)

func SaveJson(filename string, files map[string]*FileData) error {
	fmt.Printf("Writing file: %s\n", filename)
	return util.JsonWrite(filename, files)
}

func ReadJson(filename string) (map[string]*FileData, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	files := make(map[string]*FileData)
	if err := json.Unmarshal(byteValue, &files); err != nil {
		return nil, err
	}

	return files, nil
}
