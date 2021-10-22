package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readInputFiles(files StringArray) (cs []Contact, err error) {
	cs = make([]Contact, 0)
	for _, fileName := range files {
		f, err := os.Open(fileName)
		if err != nil {
			return cs, err
		}
		reader := NewLineReader(f)
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

func readStructure(lr *LineReader) ([]FieldSetter, error) {
	_, l, err := lr.ReadLine()
	if err != nil {
		return nil, err
	}

	return parseReadingTemplate(l)
}

func readContacts(lr *LineReader, setters []FieldSetter) (contacts []Contact, err error) {
	for {
		l, _, err := lr.ReadLine()
		if err == io.EOF {
			return contacts, nil
		}
		if err != nil {
			return nil, fmt.Errorf("Failed to read line %v: %w", lr.LineNumber(), err)
		}
		var fields []string
		if l != "" {
			fields = strings.Split(l, "\t")
		} else {
			continue
		}
		if len(fields) != len(setters) {
			return nil, fmt.Errorf("Line %v: Wrong number of fields.", lr.LineNumber())
		}
		contact := Contact{}
		for i, f := range fields {
			setters[i](&contact, f)
		}
		contacts = append(contacts, contact)
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
