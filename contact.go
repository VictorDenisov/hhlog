package main

import (
	"errors"
	"unicode"
)

type Frequency string

type Call string

type Date string

type Time string

type Mode string

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

func (d Date) Valid() error {
	if len(d) != 8 {
		return errors.New("Date should have six digits")
	}

	for _, c := range d {
		if !unicode.IsDigit(c) {
			return errors.New("Date should contain only digits")
		}
	}
	month := (d[4]-'0')*10 + (d[5] - '0')
	if month > 12 || month < 1 {
		return errors.New("Month should be in the range 1-12")
	}
	day := (d[6]-'0')*10 + (d[7] - '0')
	if day > 31 || day < 1 {
		return errors.New("Day should be in the range 1-31")
	}
	return nil
}

type Contact struct {
	Frequency Frequency
	Call      Call
	Date      Date
	Time      Time
	Mode      Mode
}

type FieldSetter func(c *Contact, s string)

var (
	FrequencySetter = func(c *Contact, s string) {
		c.Frequency = Frequency(s)
	}
	CallSetter = func(c *Contact, s string) {
		c.Call = Call(s)
	}
	DateSetter = func(c *Contact, s string) {
		c.Date = Date(s)
	}
	TimeSetter = func(c *Contact, s string) {
		c.Time = Time(s)
	}
	ModeSetter = func(c *Contact, s string) {
		c.Mode = Mode(s)
	}
)
