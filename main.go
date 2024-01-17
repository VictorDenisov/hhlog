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
	TSV      = "tsv"
)

func main() {
	var (
		outFormat  string
		inFile     StringArray
		template   string
		filter     string
		parkName   string
		sendPota   bool
		calcSkcc   bool
		outputFile string
	)

	flag.StringVar(&outFormat, "out", "", fmt.Sprintf("Output format: %v, %v, %v, %v", ADIF, CABRILLO, HHLOG, TSV))
	flag.StringVar(&template, "tpl", "", `Output template.

`+templateDoc())
	flag.StringVar(&outputFile, "output-file", "", "Print to file")
	flag.Var(&inFile, "in", "Input file")
	flag.StringVar(&filter, "filter", "", "Filter for QSOs")
	flag.BoolVar(&sendPota, "send-pota", false, `This flag will take the input and convert it into a file suitable for pota and wwff.

If pota section is configured in the config file it will send it to the pota coordinator. If the wwff section is configured it will also send it to wwff coordinator.

The name of the input file should have the following structure: <CALLSIGN>@<PARK>-DATE.hhl. <PARK> can be either K-<number> or KFF-<number>. The app will choose the right prefix for the submission.`)
	flag.BoolVar(&calcSkcc, "calc-skcc", false, ``)
	flag.StringVar(&parkName, "park-name", "", "Park name for pota submission")
	flag.Parse()
	filterExpr, err := ParseFilter(filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse filter expression.")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	err = checkExclusion(sendPota, calcSkcc, outFormat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	config := readConfig()
	var rawContacts []Contact

	if len(inFile) == 0 {
		rawContacts, err = readStdin()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read contacts from stdin:\n")
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	} else {
		rawContacts, err = readInputFiles(inFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read input files:\n")
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
	contacts := make([]Contact, 0)
	for _, c := range rawContacts {
		if filterExpr.run(&c) {
			contacts = append(contacts, c)
		}
	}

	if len(contacts) == 0 {
		fmt.Fprintf(os.Stderr, "No contacts parsed from input files.\n")
		flag.PrintDefaults()
		return
	}

	if sendPota {
		submitPotaReport(inFile, contacts, config, parkName)
		return
	}

	if calcSkcc {
		calcSkccScore(contacts)
		return
	}

	getters, err := parseWritingTemplate(template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse writing template: %v\n", err)
		os.Exit(1)
	}

	outF := os.Stdout
	if len(outputFile) != 0 {
		outF, err = os.Create(outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create output file: %s\n", outputFile)
			os.Exit(1)
		}
	}

	switch outFormat {
	case ADIF:
		renderAdif(outF, getters, contacts)
	case CABRILLO:
		renderCabrillo(outF, getters, contacts)
	case HHLOG:
		renderHhlog(outF, getters, contacts)
	case TSV:
		renderTsv(outF, getters, contacts)
	default:
		fmt.Fprintf(os.Stderr, "Unknown output format: %v\n", outFormat)
		fmt.Fprintf(os.Stderr, "Allowed formats are: %v, %v, %v, %v\n", ADIF, CABRILLO, HHLOG, TSV)
		os.Exit(1)
	}
}

func checkExclusion(sendPota, calcSkcc bool, outFormat string) error {
	trueCount := 0
	if outFormat != "" {
		trueCount++
	}
	if calcSkcc {
		trueCount++
	}
	if sendPota {
		trueCount++
	}
	if trueCount > 1 {
		return fmt.Errorf("Only one of send-pota, calc-skcc, out can be specified.")
	}
	return nil
}
