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

	//fmt.Println(string(b))
	// #nosec
	err = ioutil.WriteFile(filename, b, 0644)

	return err
}

//func Read(filename string) []*File {
//if b, err := json.MarshalIndent(files, "", "  "); err != nil {
//panic(err)
//} else {
//fmt.Println(string(b))
//}
//return nil
//}
