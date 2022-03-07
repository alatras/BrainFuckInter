package BrainFuckInter

import (
	"fmt"
	"testing"
)

func TestBrainFuckExecutor(t *testing.T) {
	script := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

	output := execute([]byte(script))

	fmt.Println("====", output)

	if output != "Hello World!\n" {
		t.Error("Expected output to equal 'Hello World!'\ngot", output)
	}
}

func TestValidateArguments(t *testing.T) {
	var argus []string
	argus = []string{"./bf", "--file", "./programs/sierpinskiTriangle.bf"}
	if !validateArguments(argus) {
		t.Error("Expected output to be true but got false")
	}

	argus = []string{"./bf", "--file"}
	if validateArguments(argus) {
		t.Error("Expected output to be false but got true")
	}

	argus = []string{"./bf", "--files", "./programs/sierpinskiTriangle.bf"}
	if validateArguments(argus) {
		t.Error("Expected output to be false but got true")
	}
}
