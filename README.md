# Brainfuck Interpreter

Golang library to interpret BrainFrack code.

## Get

**Download:**

```
go get github.com/alatras/BrainFuckInter
```

**Import:**

```
import "github.com/alatras/BrainFuckInter"
```

## Use

#### 1- Live interpret your input:

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

The above operations also apply using bellow method `InterpretOnce` if letters are added to the script.

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

Where `os.Argus` are three arguments that are required: `[your Go app] --file/--script filename/Brainfuck script`

**Examples:**

1- "Hello World!" with Brainfuck script as an argument:

```
[your Go app] --script  "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
```

2- "Hello World!" with a file containing Brainfuck script:

```
[your Go app] --file ./helloworld.bf
```

## Test

```
go test -v ./...
```

**Note**: You can download 2 examples of Brainfuck programs from [brainfuck-programs](https://github.com/alatras/brainfuck-programs) to test the method `InterpretOnce` with `--file`.
