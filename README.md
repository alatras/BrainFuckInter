# Brainfuck Interpreter

Golang library to interpret Brainfuck code.

## Download

**Get:**

```
go get github.com/alatras/BrainFuckInter
```

**Import:**

```
import "github.com/alatras/BrainFuckInter"
```

## Use

#### 1- Live-interpret your input:

```
func main() {
	err := BrainFuckInter.InterpretLive()
	if err != nil {
		log.Fatal(err)
	}
}
```

Run your app and you will be prompted to start typing your Brainfuck script. You get the output in console after every entry. You get `¯\_(ツ)_/¯` when there is nothing show.

**Custom operations at runtime:**

- Type `s` to square current cell value.
- Type `h` to half current cell value.
- Type `c` to cube current cell value.


#### 2- Interpret a script passed as an argument or stored in file:

```
func main() {
	output, err := BrainFuckInter.InterpretOnce(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
```

Where `os.Argus` are three arguments that **are required**. Syntax:
`[your Go app] --file/--script filename/Brainfuck-script`

The **custom operations** mentioned above also apply using this method if corresponding letters (`s, h, c`) are added to the script.

**Examples:**

1- "Hello World!" with Brainfuck script as an argument:

```
[your Go app] --script  "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
```

2- "Hello World!" with a file containing Brainfuck script:

```
[your Go app] --file ./helloworld.bf
```

You can download examples of Brainfuck programs from [brainfuck-programs](https://github.com/alatras/brainfuck-programs) to test this method with `--file`.

## Run tests

```
go test -v ./...
```
