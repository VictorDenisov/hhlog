package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

const (
	FREQUENCY = "%f"
	CALL      = "%c"
	DATE      = "%d"
	TIME      = "%t"
	BAND      = "%b"
	MODE      = "%m"
)

func readInputFiles(files StringArray) (cs []Contact, err error) {
	cs = make([]Contact, 0)
	for _, fileName := range files {
		f, err := os.Open(fileName)
		if err != nil {
			return cs, err
		}
		reader := bufio.NewReader(f)
		setters, err := readStructure(reader)
		if err != nil {
			return nil, err
		}
		contacts, err := readContacts(reader, setters)
		if err != nil {
			return nil, err
		}
		cs = append(cs, contacts...)
	}
	return cs, nil
}

func readStructure(reader *bufio.Reader) ([]FieldSetter, error) {
	line := ""
	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		trimmedLine := strings.TrimSpace(string(l))
		if trimmedLine != "" {
			line = strings.TrimSpace(trimmedLine[1:])
			break
		}
	}

	verbs := strings.Split(line, "\t")
	setters := make([]FieldSetter, len(verbs))
	for i, v := range verbs {
		switch v {
		case FREQUENCY:
			setters[i] = FrequencySetter
		case CALL:
			setters[i] = CallSetter
		case DATE:
			setters[i] = DateSetter
		case TIME:
			setters[i] = TimeSetter
		case MODE:
			setters[i] = ModeSetter
		}
	}
	return setters, nil
}

func readContacts(reader *bufio.Reader, setters []FieldSetter) (contacts []Contact, err error) {
	for {
		l, _, err := reader.ReadLine()
		if err == io.EOF {
			return contacts, nil
		}
		if err != nil {
			return nil, err
		}
		trimmedLine := strings.TrimSpace(string(l))
		var fields []string
		if trimmedLine != "" {
			fields = strings.Split(trimmedLine, "\t")
		}
		if len(fields) != len(setters) {
			return nil, errors.New("The number of fields in a line doesn't match the template")
		}
		contact := Contact{}
		for i, f := range fields {
			setters[i](&contact, f)
		}
		contacts = append(contacts, contact)
	}
}
