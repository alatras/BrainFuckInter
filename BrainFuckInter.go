// BrainFuckInter package
//
// The main method, Interpret, receives client arguments consisting of:
// 1- App launch command
// 2- Type of argument (--file or --script)
// 3- File path or script
//
// Responses with Brainfuck output logged into the console.

package BrainFuckInter

import (
	"errors"
)

//
// Send file contents (should be Brainfuck script) to executer.
//
func interpretFile(args []string) (string, error) {
	contents, err := readFile(args[2])
	if err != nil {
		return "", errors.New(err.Error())
	}
	return execute(contents), nil
}

//
// Send script argument (should be Brainfuck script) to executer.
//
func interpretScript(args []string) (string, error) {
	return execute([]byte(args[2])), nil
}

//
// Main method.
// It validates, store commands and assing arguments to executor.
//
func Interpret(args []string) (string, error) {
	if !validateArguments(args) {
		return "", errors.New("syntax: '[app] --file/--script filename/Brainfuck script'")
	}

	err := addCommandToHistory(args)
	if err != nil {
		return "", errors.New("error writing history:" + err.Error())
	}

	switch args[1] {
	case "--script":
		return interpretScript(args)
	case "--file":
		return interpretFile(args)
	default:
		return "", errors.New("unknown error")
	}
}
