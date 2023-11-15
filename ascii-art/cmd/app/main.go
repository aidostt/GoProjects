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
	err = validator(input, desiredFont, flags)
	if err != nil {
		fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
		fmt.Println(err)
		return
	}

	//TODO:implement the colors

	alphabet, err := internal.Alphabet(desiredFont, lettersToColorize)

	if err != nil {
		fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
		fmt.Println(err)
		return
	}
	//TODO:implement the justify
	out := internal.FormatOutput(alphabet, input)
	if output != "" {
		file, err := pkg.File(output)
		defer file.Close()
		if err != nil {
			fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
			fmt.Println(err)
			return
		}
		err = pkg.PrintFile(file, out)
		if err != nil {
			fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println(out)
	}
}
