package goma

import (
	"io/ioutil"
	"fmt"
)

// TODO cache
func get(path string) (data string, err error) {
	result := []byte(nil)
	result, err = ioutil.ReadFile(path)
	return fmt.Sprintf("%s", result), err
}

// replace named key