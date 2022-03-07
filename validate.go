package BrainFuckInter

import (
	"os"

	"github.com/fatih/color"
)

func validate(args []string) {
	if len(args) != 2 {
		color.Yellow("Use 'bf filename'\n")
		os.Exit(0)
	}
}
