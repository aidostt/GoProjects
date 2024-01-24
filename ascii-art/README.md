# ascii-art

This is a command-line program written in Go that takes a string as an argument and outputs a graphic representation of the string using ASCII characters. The program uses various ASCII art templates, including shadow, standard, and thinkertoy.

## Usage
Go to the project directory

```bash
cd /cmd/main/
```


To use the ASCII Art Generator, you can run the program with the desired string as an argument:

```bash
go run . "Your text goes here"
```

The program will then display the graphic representation of the input string using ASCII characters.
Examples

Here are some examples of running the program with different input strings:

```bash
go run . "" | cat -e
go run . "\n" | cat -e
go run . "Hello\n" | cat -e
go run . "hello" | cat -e
go run . "HeLlO" | cat -e
go run . "Hello There" | cat -e
go run . "1Hello 2There" | cat -e
go run . "{Hello There}" | cat -e
go run . "Hello\nThere" | cat -e
go run . "Hello\n\nThere" | cat -e
```

# Banner Format

The program uses predefined ASCII art templates for each character. Each character is represented by a 8-line high banner, and characters are separated by a new line (\n) in the templates.

# Contributing

Contributions to this project are welcome. If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.
