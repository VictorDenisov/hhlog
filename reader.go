package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readStdin() (cs []Contact, err error) {
	lineReader := NewLineReader(os.Stdin)
	contacts, err := readContacts(lineReader)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func readInputFiles(files StringArray) (cs []Contact, err error) {
	cs = make([]Contact, 0)
	for _, fileName := range files {
		f, err := os.Open(fileName)
		if err != nil {
			return cs, err
		}
		lineReader := NewLineReader(f)
		contacts, err := readContacts(lineReader)
		if err != nil {
			return nil, err
		}
		cs = append(cs, contacts...)
	}
	return cs, nil
}

func readStructure(lr *LineReader) ([]FieldSetter, error) {
	_, c, err := lr.ReadLine()
	if err != nil {
		return nil, err
	}

	return parseReadingTemplate(c)
}

func readContacts(lr *LineReader) (contacts []Contact, err error) {
	var setters []FieldSetter
	contact := Contact{}
	for {
		l, c, err := lr.ReadLine()
		if err == io.EOF {
			return contacts, nil
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to read line %v: %w", lr.LineNumber(), err)
		}
		if len(l) == 0 { // Probably new template string
			if isTemplateString(c) {
				setters, err = parseReadingTemplate(c)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Attempted to parse template from line: %v\n", lr.LineNumber())
					fmt.Fprintf(os.Stderr, "Failed parsing with error: %v\n", err)
					fmt.Fprintf(os.Stderr, "Continuing processing with old template.\n")
				}
			}
		} else {
			var fields []string
			fields = strings.Split(l, "\t")
			if len(fields) > len(setters) {
				return nil, fmt.Errorf("Line %v: Contains more values than fields.", lr.LineNumber())
			} // Less fields is fine. The missing fields will be taken from the most recent contact with this value specified.
			for i, f := range fields {
				setters[i](&contact, f)
			}
			contacts = append(contacts, contact)
		}
	}
}

type LineReader struct {
	reader     *bufio.Reader
	lineNumber int
}

func NewLineReader(f io.Reader) *LineReader {
	return &LineReader{bufio.NewReader(f), 0}
}

func (lr *LineReader) ReadLine() (line string, comment string, err error) {
	lr.lineNumber++
	l, err := lr.reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	p := strings.Index(l, "\"")
	if p == -1 {
		return strings.TrimSpace(l), "", nil
	}
	return strings.TrimSpace(l[0:p]), strings.TrimSpace(l[p+1 : len(l)]), nil
}

func (lr *LineReader) LineNumber() int {
	return lr.lineNumber
}
