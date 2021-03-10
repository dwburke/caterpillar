package util

import (
	"encoding/json"
	"io/ioutil"
)

func JsonWrite(filename string, data interface{}) error {

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// #nosec
	err = ioutil.WriteFile(filename, b, 0644)

	return err
}
