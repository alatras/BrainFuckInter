package BrainFuckInter

import (
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

func readFile(args []string) []byte {
	filename := args[1]

	program, err := ioutil.ReadFile(filename)
	if err != nil {
		color.Red("Cannot read/find file: '%s'\n", filename)
		os.Exit(0)
	}

	return program
}
