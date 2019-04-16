package goma

import (
	"fmt"
	"io/ioutil"
)

// TODO cache
func getFile(path string) (data string, err error) {
	result := []byte(nil)
	result, err = ioutil.ReadFile(path)
	return fmt.Sprintf("%s", result), err
}

// replace named key
