package main

import (
	"errors"
	"strings"
	"unicode"
)

type Frequency string
type Call string
type Date string
type Time string
type Mode string
type Srx string

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
	Srx       Srx
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
	SrxSetter = func(c *Contact, s string) {
		c.Srx = Srx(s)
	}
)

type FieldGetterVisitor interface {
	visitFrequency(g *FrequencyGetter)
	visitCall(g *CallGetter)
	visitDate(g *DateGetter)
	visitTime(g *TimeGetter)
	visitMode(g *ModeGetter)
	visitBand(g *BandGetter)
	visitSkcc(g *SkccGetter)
	visitName(g *NameGetter)
	visitSpc(g *SpcGetter)
}

type FieldGetter interface {
	get(c *Contact)
	accept(v FieldGetterVisitor)
}

type FrequencyGetter struct {
	val string
}

func (g *FrequencyGetter) get(c *Contact) {
	g.val = string(c.Frequency)
}

func (g *FrequencyGetter) accept(v FieldGetterVisitor) {
	v.visitFrequency(g)
}

type CallGetter struct {
	val string
}

func (g *CallGetter) get(c *Contact) {
	g.val = string(c.Call)
}

func (g *CallGetter) accept(v FieldGetterVisitor) {
	v.visitCall(g)
}

type DateGetter struct {
	val string
}

func (g *DateGetter) get(c *Contact) {
	g.val = string(c.Date)
}

func (g *DateGetter) accept(v FieldGetterVisitor) {
	v.visitDate(g)
}

type TimeGetter struct {
	val string
}

func (g *TimeGetter) get(c *Contact) {
	g.val = string(c.Time)
}

func (g *TimeGetter) accept(v FieldGetterVisitor) {
	v.visitTime(g)
}

type ModeGetter struct {
	val string
}

func (g *ModeGetter) get(c *Contact) {
	g.val = string(c.Mode)
}

func (g *ModeGetter) accept(v FieldGetterVisitor) {
	v.visitMode(g)
}

type BandGetter struct {
	val string
}

func (g *BandGetter) get(c *Contact) {
	f := string(c.Frequency)
	if strings.HasPrefix(f, "7") {
		g.val = "40M"
	}
	if strings.HasPrefix(f, "14") {
		g.val = "20M"
	}
}

func (g *BandGetter) accept(v FieldGetterVisitor) {
	v.visitBand(g)
}

type SkccGetter struct {
	db  *SkccDB
	val string
}

func (g *SkccGetter) get(c *Contact) {
	call := string(c.Call)
	g.val = g.db.callIndex[call].Skcc
}

func (g *SkccGetter) accept(v FieldGetterVisitor) {
	v.visitSkcc(g)
}

type NameGetter struct {
	db  *SkccDB
	val string
}

func (g *NameGetter) get(c *Contact) {
	call := string(c.Call)
	g.val = g.db.callIndex[call].Name
}

func (g *NameGetter) accept(v FieldGetterVisitor) {
	v.visitName(g)
}

type SpcGetter struct {
	db  *SkccDB
	val string
}

func (g *SpcGetter) get(c *Contact) {
	call := string(c.Call)
	g.val = g.db.callIndex[call].Spc
}

func (g *SpcGetter) accept(v FieldGetterVisitor) {
	v.visitSpc(g)
}
