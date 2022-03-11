# Brainfuck Interpreter

Go library to interpret Brainfuck code.

## Download

```bash
$ go get github.com/alatras/BrainFuckInter
```

## Use

#### 1- Live interpreter

Use `InterpretLive` for getting live interpretation while typing Brainfuck code.

```go
import "github.com/alatras/BrainFuckInter"

func main() {
	err := BrainFuckInter.InterpretLive()
	if err != nil {
		log.Fatal(err)
	}
}
```

Run your app and you will be prompted to start typing your Brainfuck script. You get the output in console after every entry. When there is nothing to show, you get `¯\_(ツ)_/¯` only.

**Custom operations at runtime:**

- Type `s` to square current cell value.
- Type `h` to half current cell value.
- Type `c` to cube current cell value.


#### 2- One-time interpreter

Use `InterpretOnce` for one-time interpretation by passing ready Brainfuck code as an argument or file.

```go
import "github.com/alatras/BrainFuckInter"

func main() {
	output, err := BrainFuckInter.InterpretOnce(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
```

Where `os.Argus` are three arguments that **are required**. Syntax:
`[your Go app] --file/--script filename/Brainfuck script`

**Examples:**

1- Pass "Hello World!" code as an argument:

```bash
$ [your Go app] --script  "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
```

2- Pass "Hello World!" code from a file containing code:

```bash
$ [your Go app] --file ./helloworld.bf
```

You can get examples of Brainfuck programs from [brainfuck-programs](https://github.com/alatras/brainfuck-programs) to test this method with `--file`.

The **custom operations** added to above method (`InterpretLive`) also apply in this method if corresponding letters (`s, h, c`) are added to passed Brainfuck code.

## Run tests

```bash
$ go test -v ./...
```
