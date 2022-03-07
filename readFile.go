package BrainFuckInter

import (
	"errors"
	"io/ioutil"
)

/*
readFile:
Read and return fill contents given its path.
*/
func readFile(path string) ([]byte, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("Cannot read/find file:" + path)
	}

	return contents, nil
}
