/*
/ BrainFuckInter package
*/

package BrainFuckInter

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

/*
/ interpretFile sends file contents (should be Brainfuck script) to executer.
*/
func interpretFile(filePath string) (string, error) {
	contents, err := readFile(filePath)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return execute(contents), nil
}

/*
/ interpretScript sends script argument (should be Brainfuck script) to executer.
*/
func interpretScript(script string) (string, error) {
	return execute([]byte(script)), nil
}

/*
/ InterpretOnce interprets from a set of arguments entered one time,
/ returns the results,
/ validate arguments,
/ and stores commands in history.
*/
func InterpretOnce(args []string) (string, error) {
	if !validateArguments(args) {
		return "", errors.New("syntax: '[app] --file/--script filename/Brainfuck script'")
	}

	err := addCommandToHistory(args)
	if err != nil {
		return "", errors.New("error writing history:" + err.Error())
	}

	switch args[1] {
	case "--script":
		return interpretScript(args[2])
	case "--file":
		return interpretFile(args[2])
	default:
		return "", errors.New("unknown error")
	}
}

/*
/ InterpretLive interprets user input character after another,
/ validates, stores commands and assigns arguments to executor.
*/
func InterpretLive() error {
	color.Cyan("Start typing Brainfuck script: ")

	script := []string{}
	input := ""
	for {
		fmt.Scanln(&input)
		if input != "" {
			script = append(script, input)

			output, err := interpretScript(strings.Join(script, ""))
			if err != nil {
				return errors.New(err.Error())
			}

			if output == "" {
				color.Yellow("¯\\_(ツ)_/¯")
			} else {
				color.Green(output)
			}
		}
		input = ""
	}
}
