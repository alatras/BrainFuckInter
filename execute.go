package BrainFuckInter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
/ execute executes Brainfuck logic.
*/
func execute(program []byte) string {
	input := bufio.NewReader(os.Stdin)
	stack := make([]byte, 30720)
	loops := make([]int, 0, 1024)
	cell := 0
	skip := 0

	var output strings.Builder

	for x := 0; x < len(program); x++ {

		switch program[x] {

		case '>':
			cell++

		case '<':
			cell--

		case '+':
			stack[cell]++

		case 's':
			stack[cell] = stack[cell] * stack[cell]

		case 'h':
			stack[cell] = stack[cell] / 2

		case 'c':
			stack[cell] = stack[cell] * stack[cell] * stack[cell]

		case '-':
			stack[cell]--

		case '.':
			bar := fmt.Sprintf("%c", stack[cell])
			output.WriteString(bar)

		case ',':
			stack[cell], _ = input.ReadByte()

		case '[':
			if stack[cell] == 0 {
				skip++
				for skip > 0 {
					x++
					if program[x] == '[' {
						skip++
					} else if program[x] == ']' {
						skip--
					}
				}
			} else {
				loops = append(loops, x)
			}

		case ']':
			if stack[cell] == 0 {
				loops = loops[:len(loops)-1]
			} else {
				x = loops[len(loops)-1]
			}
		}
	}

	return output.String()
}
