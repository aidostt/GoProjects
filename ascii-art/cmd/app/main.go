package main

import (
	"aidostt.ascii-art/internal"
	"aidostt.ascii-art/pkg"
	"flag"
	"fmt"
)

func main() {
	var (
		color, align, reverse, output string
	)
	flag.StringVar(&output, "output", "", "file name for writing results in it")
	flag.StringVar(&reverse, "reverse", "", "file name for taking data to reverse")
	flag.StringVar(&align, "align", "left", "justify text")
	flag.StringVar(&color, "color", "", "letter color")
	flag.Parse()
	args := flag.Args()
	flags := map[string]string{
		"output":  output,
		"reverse": reverse,
		"align":   align,
		"color":   color,
	}
	input, lettersToColorize, desiredFont, err := assignArgs(args)
	if err != nil {
		fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
		fmt.Println(err)
		return
	}
	fmt.Println(lettersToColorize)
	err = validator(input, desiredFont, flags)
	if err != nil {
		fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
		fmt.Println(err)
		return
	}

	//TODO:implement the colors

	alphabet, err := internal.Alphabet(desiredFont)

	if err != nil {
		fmt.Println(err)
		return
	}
	//TODO:implement the justify
	out := internal.FormatOutput(alphabet, input)
	//TODO:implement the output
	fmt.Println(out)
}
