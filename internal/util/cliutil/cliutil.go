package cliutil

import (
	"io/ioutil"
	"os"
)

func ReadFromFileOrStdin(filename string) ([]byte, error) {
	if filename != "" {
		return ioutil.ReadFile(filename)
	}
	return ioutil.ReadAll(os.Stdin)
}
