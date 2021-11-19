package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ADIF     = "adi"
	CABRILLO = "cbr"
	HHLOG    = "hhl"
)

func main() {
	var (
		outFormat string
		inFile    StringArray
		template  string
	)

	flag.StringVar(&outFormat, "out", "", fmt.Sprintf("Output format: %v, %v, %v", ADIF, CABRILLO, HHLOG))
	flag.StringVar(&template, "tpl", "", `Output template.

`+templateDoc())
	flag.Var(&inFile, "in", "Input file")
	flag.Parse()

	contacts, err := readInputFiles(inFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input files:\n")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if len(contacts) == 0 {
		fmt.Fprintf(os.Stderr, "No contacts parsed from input files.\n")
		flag.PrintDefaults()
		return
	}

	getters, err := parseWritingTemplate(template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse writing template: %v\n", err)
		os.Exit(1)
	}

	switch outFormat {
	case ADIF:
		renderAdif(getters, contacts)
	case CABRILLO:
		renderCabrillo(getters, contacts)
	default:
		fmt.Fprintf(os.Stderr, "Unknown output format: %v\n", outFormat)
		fmt.Fprintf(os.Stderr, "Allowed formats are: %v, %v\n", ADIF, CABRILLO)
		os.Exit(1)
	}
}
