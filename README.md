# Brainfuck Interpreter

Simple Go library that interprets BrainFrack code.

## Usage

**Download:**

```
go get github.com/alatras/BrainFuckInter
```

**Import:**

```
import "github.com/alatras/BrainFuckInter"
```

## Code

#### 1- Live interpret your input:

```
func main() {
	err := BrainFuckInter.InterpretLive()
	if err != nil {
		log.Fatal(err)
	}
}
```

Then you can start entering your Brainfuck script. You get the output after every new line.

**Extra functionalities on the fly:**

- Type `s` to square current cell value.
- Type `h` to half current cell value.
- Type `c` to cube current cell value.

#### 2- Interpret a script entered all at once or stored in a file:

```
func main() {
	output, err := BrainFuckInter.InterpretOnce(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
```

Where `os.Argus` are three arguments that are required: `[app] --file/--script filename/Brainfuck script`

Examples:

1- "Hello World!" with Brainfuck script as an argument:

```
go run *.go --script  "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
```

2- "Hello World!" with a file containing Brainfuck script:

```
go run *.go --file ./helloworld.bf
```

## Test

```
go test -v ./...
```

**Note**: The repo [brainfuck-programs](https://github.com/alatras/brainfuck-programs) includes two examples of Brainfuck programs. You can download them to test `InterpretOnce` method with `--file`.
