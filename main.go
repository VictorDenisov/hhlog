package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ADIF     = "adi"
	CABRILLO = "cbr"
)

func main() {
	var (
		outFormat string
		inFile    StringArray
		template  string
	)

	flag.StringVar(&outFormat, "out", "", "Output format")
	flag.StringVar(&template, "tpl", "", "Output template")
	flag.Var(&inFile, "in", "Input file")
	flag.Parse()

	contacts, err := readInputFiles(inFile)
	if err != nil {
		fmt.Printf("Failed to read input files\n")
		os.Exit(1)
	}

	switch outFormat {
	case ADIF:
		fmt.Printf("%v\n", contacts)
	case CABRILLO:
	default:
		fmt.Printf("Unknown output format: %v\n", outFormat)
		fmt.Printf("Allowed formats are: %v, %v\n", ADIF, CABRILLO)
		os.Exit(1)
	}
}
