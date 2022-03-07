package BrainFuckInter

import (
	"os"
)

func Interpret() {
	args := os.Args

	validate(args)

	program := readFile(args)

	execute(program)
}
