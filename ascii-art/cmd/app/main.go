package main

import (
	"aidostt.ascii-art/internal"
	"aidostt.ascii-art/pkg"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var (
		colorFlag, alignFlag, reverseFlag, outputFlag string
	)
	flag.StringVar(&outputFlag, "output", "", "file name for writing results in it")
	flag.StringVar(&reverseFlag, "reverse", "", "file name for taking data to reverseFlag")
	flag.StringVar(&alignFlag, "align", "", "justify text")
	flag.StringVar(&colorFlag, "color", "", "letter colorFlag")
	flag.Parse()
	args := flag.Args()
	flags := map[string]string{
		"output":  outputFlag,
		"reverse": reverseFlag,
		"align":   alignFlag,
		"color":   colorFlag,
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
	alphabet, err := internal.Alphabet(desiredFont)

	if err != nil {
		fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
		fmt.Println(err)
		return
	}
	output := internal.FormatOutput(alphabet, input, lettersToColorize, strings.ToLower(flags["color"]))
	if flags["align"] != "" {
		output, err = internal.Justify(flags["align"], output)
		if err != nil {
			fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
			fmt.Println(err)
			return
		}
	}
	if flags["output"] != "" {
		file, err := pkg.File(flags["output"])
		defer file.Close()
		if err != nil {
			fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
			fmt.Println(err)
			return
		}
		err = pkg.PrintFile(file, output)
		if err != nil {
			fmt.Printf("Occured error: %v\n", pkg.ErrInvalidInput)
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println(output)
	}
}
