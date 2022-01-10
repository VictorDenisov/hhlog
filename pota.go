package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

func submitPotaReport(inFile []string, contacts []Contact, config *Config, parkName string) error {
	if len(inFile) != 1 {
		return fmt.Errorf("Submit-pota command requires exactly one input file.")
	}
	if len(parkName) == 0 {
		return fmt.Errorf("Park name is required for pota submissions.")
	}
	getters, err := parseWritingTemplate("%c %b %m %d %t %f")
	if err != nil {
		panic("Failed to parse predefined output template")
	}

	callSign, parkCode, date := parseInputFile(inFile[0])

	potaFileName := fmt.Sprintf("%v@K-%v-%v.adi", callSign, parkCode, date)
	wwffFileName := fmt.Sprintf("%v@KFF-%v-%v.adi", callSign, parkCode, date)

	f, err := os.Create(potaFileName)
	if err != nil {
		return err
	}
	renderAdif(f, getters, contacts)
	err = f.Close()
	if err != nil {
		return err
	}

	d := gomail.NewDialer(config.Station.Mail.SmtpHost, config.Station.Mail.Port, config.Station.Mail.Email, config.Station.Mail.Password)
	if config.Pota.ContactEmail != "" && config.Pota.ContactName != "" {
		m := gomail.NewMessage()
		m.SetHeader("From", config.Station.Mail.Email)
		m.SetHeader("To", config.Pota.ContactEmail)
		m.SetHeader("Subject", fmt.Sprintf("%v K-%v %s", callSign, parkCode, parkName))
		m.SetBody("text/plain", fmt.Sprintf("Hello %s,\n\nHere is my log for K-%v on %v.\n\nThanks, Victor.", config.Pota.ContactName, parkCode, date))
		m.Attach(potaFileName)

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
	if config.Wwff.ContactEmail != "" && config.Wwff.ContactName != "" {
		m := gomail.NewMessage()
		m.SetHeader("From", config.Station.Mail.Email)
		m.SetHeader("To", config.Wwff.ContactEmail)
		m.SetHeader("Subject", fmt.Sprintf("%v KFF-%v %s", callSign, parkCode, parkName))
		m.SetBody("text/plain", fmt.Sprintf("Hello %s,\n\nHere is my log for KFF-%v on %v.\n\nThanks, Victor.", config.Wwff.ContactName, parkCode, date))
		m.Attach(potaFileName, gomail.Rename(wwffFileName))

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
	return nil
}

func parseInputFile(s string) (Call, string, Date) {
	at := strings.Index(s, "@")
	dash := strings.Index(s, "-")
	dash2 := strings.Index(s[dash+1:], "-") + dash + 1
	dot := strings.Index(s, ".")
	return Call(s[0:at]), s[dash+1 : dash2], Date(s[dash2+1 : dot])
}
