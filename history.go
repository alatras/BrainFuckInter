package BrainFuckInter

import (
	"os"
	"strings"
)

const historyFile = "history.txt"

/*
addCommandToHistory:
Store commands in history file.
*/
func addCommandToHistory(args []string) error {
	f, err := os.OpenFile(historyFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(strings.Join(args, " ") + "\n"); err != nil {
		return err
	}

	return nil
}
