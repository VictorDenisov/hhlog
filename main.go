package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ADIF     = "adi"
	CABRILLO = "cbr"

	FREQUENCY = "%f"
	CALL      = "%c"
	DATE      = "%d"
	TIME      = "%t"
	BAND      = "%b"
	MODE      = "%m"
)

func main() {
	var (
		outFormat string
		inFile    StringArray
	)

	flag.StringVar(&outFormat, "out", "", "Output format")
	flag.Var(&inFile, "in", "Input file")
	flag.Parse()

	switch outFormat {
	case ADIF:
		fmt.Printf("%v\n", inFile)
	case CABRILLO:
	default:
		fmt.Printf("Unknown output format: %v\n", outFormat)
		fmt.Printf("Allowed formats are: %v, %v\n", ADIF, CABRILLO)
		os.Exit(1)
	}
}
