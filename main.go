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
	flag.StringVar(&template, "tpl", "", `Output template.

%f - frequency in megahertz
%c - call sign
%d - eight digits of date: year month day
%t - four digits of UTC time
%b - band. Only used in output template
%m - mode
`)
	flag.Var(&inFile, "in", "Input file")
	flag.Parse()

	contacts, err := readInputFiles(inFile)
	if err != nil {
		fmt.Printf("Failed to read input files:\n")
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	if len(contacts) == 0 {
		fmt.Printf("No contacts parsed from input files.\n")
		flag.PrintDefaults()
		return
	}

	getters, err := parseWritingTemplate(template)
	if err != nil {
		fmt.Printf("Failed to parse writing template: %v\n", err)
		os.Exit(1)
	}

	switch outFormat {
	case ADIF:
		renderAdif(getters, contacts)
	case CABRILLO:
	default:
		fmt.Printf("Unknown output format: %v\n", outFormat)
		fmt.Printf("Allowed formats are: %v, %v\n", ADIF, CABRILLO)
		os.Exit(1)
	}
}
