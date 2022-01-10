package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
		filter    string
		sendPota  bool
	)

	flag.StringVar(&outFormat, "out", "", fmt.Sprintf("Output format: %v, %v, %v", ADIF, CABRILLO, HHLOG))
	flag.StringVar(&template, "tpl", "", `Output template.

`+templateDoc())
	flag.Var(&inFile, "in", "Input file")
	flag.StringVar(&filter, "filter", "", "Filter for QSOs")
	flag.BoolVar(&sendPota, "send-pota", false, `This flag will take the input and convert it into a file suitable for pota and wwff.

If pota section is configured in the config file it will send it to the pota coordinator. If the wwff section is configured it will also send it to wwff coordinator.

The name of the input file should have the following structure: <CALLSIGN>@<PARK>-DATE.hhl. <PARK> can be either K-<number> or KFF-<number>. The app will choose the right prefix for the submission.`)
	flag.Parse()
	filterExpr, err := ParseFilter(filter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse filter expression.")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	var config *Config
	data, err := ioutil.ReadFile(".hhlog.conf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse config file:\n")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "Proceeding without config file.\n")
	} else {
		config, err = parseConfig(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse config file:\n")
			fmt.Fprintf(os.Stderr, "%v\n", err)
			fmt.Fprintf(os.Stderr, "Proceeding without config file.\n")
			config = nil
		}
	}

	rawContacts, err := readInputFiles(inFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input files:\n")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
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
		submitPotaReport(inFile, contacts, config)
		return
	}

	getters, err := parseWritingTemplate(template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse writing template: %v\n", err)
		os.Exit(1)
	}

	switch outFormat {
	case ADIF:
		renderAdif(os.Stdout, getters, contacts)
	case CABRILLO:
		renderCabrillo(getters, contacts)
	case HHLOG:
		renderHhlog(os.Stdout, getters, contacts)
	default:
		fmt.Fprintf(os.Stderr, "Unknown output format: %v\n", outFormat)
		fmt.Fprintf(os.Stderr, "Allowed formats are: %v, %v\n", ADIF, CABRILLO, HHLOG)
		os.Exit(1)
	}
}
