package main

import (
	"errors"
	"unicode"
)

type Frequency string
type Call string

type Date string

type Time string

func (t Time) Valid() error {
	if len(t) != 4 {
		return errors.New("Time should have four digits")
	}
	for _, c := range t {
		if !unicode.IsDigit(c) {
			return errors.New("Time should contain only digits")
		}
	}
	if (t[0]-'0')*10+(t[1]-'0') > 23 {
		return errors.New("Hour should be in the range 00-23")
	}
	if (t[2]-'0')*10+(t[3]-'0') > 59 {
		return errors.New("Minutes should be in the range 00-59")
	}
	return nil
}

type Contact struct {
	Frequency string
	Call      string
	Date      string
	Time      string
}
