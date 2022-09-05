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
type Skcc string
type Srx string
type Stx string
type Prec string
type Ck string
type Sect string
type Spc string
type RstRcvd string
type RstSent string
type State string
type MySotaRef string
type MyState string
type Cnty string

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

// Implementation of read contacts relies on the fact that fields of this
// structure are immutable.
type Contact struct {
	Frequency Frequency
	Call      Call
	Date      Date
	Time      Time
	Mode      Mode
	Skcc      Skcc
	Spc       Spc
	Srx       Srx
	Stx       Stx
	Prec      Prec
	Ck        Ck
	Sect      Sect
	RstRcvd   RstRcvd
	RstSent   RstSent
	State     State
	MySotaRef MySotaRef
	MyState   MyState
	Cnty      Cnty
}

type FieldSetter func(c *Contact, s string)

type FieldSetterConstructor func() FieldSetter

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
	SkccSetter = func(c *Contact, s string) {
		c.Skcc = Skcc(s)
	}
	SpcSetter = func(c *Contact, s string) {
		c.Spc = Spc(s)
	}
	SrxSetter = func(c *Contact, s string) {
		c.Srx = Srx(s)
	}
	StxSetter = func(c *Contact, s string) {
		c.Stx = Stx(s)
	}
	PrecSetter = func(c *Contact, s string) {
		c.Prec = Prec(s)
	}
	CkSetter = func(c *Contact, s string) {
		c.Ck = Ck(s)
	}
	SectSetter = func(c *Contact, s string) {
		c.Sect = Sect(s)
	}
	RstRcvdSetter = func(c *Contact, s string) {
		c.RstRcvd = RstRcvd(s)
	}
	RstSentSetter = func(c *Contact, s string) {
		c.RstSent = RstSent(s)
	}
	StateSetter = func(c *Contact, s string) {
		c.State = State(s)
	}
	MySotaRefSetter = func(c *Contact, s string) {
		c.MySotaRef = MySotaRef(s)
	}
	MyStateSetter = func(c *Contact, s string) {
		c.MyState = MyState(s)
	}
	CntySetter = func(c *Contact, s string) {
		c.Cnty = Cnty(s)
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
	visitSrx(g *SrxGetter)
	visitStx(g *StxGetter)
	visitPrec(g *PrecGetter)
	visitCk(g *CkGetter)
	visitSect(g *SectGetter)
	visitRstRcvd(g *RstRcvdGetter)
	visitRstSent(g *RstSentGetter)
	visitState(g *StateGetter)
	visitMySotaRef(g *MySotaRefGetter)
	visitMyState(g *MyStateGetter)
	visitCnty(g *CntyGetter)
}

type FieldGetter interface {
	get(c *Contact)
	accept(v FieldGetterVisitor)
}

type FieldGetterConstructor func() FieldGetter

type FrequencyGetter struct {
	val Frequency
}

func (g *FrequencyGetter) get(c *Contact) {
	g.val = c.Frequency
}

func (g *FrequencyGetter) accept(v FieldGetterVisitor) {
	v.visitFrequency(g)
}

type CallGetter struct {
	val Call
}

func (g *CallGetter) get(c *Contact) {
	g.val = c.Call
}

func (g *CallGetter) accept(v FieldGetterVisitor) {
	v.visitCall(g)
}

type DateGetter struct {
	val Date
}

func (g *DateGetter) get(c *Contact) {
	g.val = c.Date
}

func (g *DateGetter) accept(v FieldGetterVisitor) {
	v.visitDate(g)
}

type TimeGetter struct {
	val Time
}

func (g *TimeGetter) get(c *Contact) {
	g.val = c.Time
}

func (g *TimeGetter) accept(v FieldGetterVisitor) {
	v.visitTime(g)
}

type ModeGetter struct {
	val Mode
}

func (g *ModeGetter) get(c *Contact) {
	g.val = c.Mode
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
	val Skcc
}

func (g *SkccGetter) get(c *Contact) {
	if c.Skcc == "" {
		call := string(c.Call)
		g.val = g.db.callIndex[call].Skcc
	} else {
		g.val = c.Skcc
	}
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
	val Spc
}

func (g *SpcGetter) get(c *Contact) {
	if c.Spc == "" {
		call := string(c.Call)
		g.val = g.db.callIndex[call].Spc
	} else {
		g.val = c.Spc
	}
}

func (g *SpcGetter) accept(v FieldGetterVisitor) {
	v.visitSpc(g)
}

type SrxGetter struct {
	val Srx
}

func (g *SrxGetter) get(c *Contact) {
	g.val = c.Srx
}

func (g *SrxGetter) accept(v FieldGetterVisitor) {
	v.visitSrx(g)
}

type StxGetter struct {
	val Stx
}

func (g *StxGetter) get(c *Contact) {
	g.val = c.Stx
}

func (g *StxGetter) accept(v FieldGetterVisitor) {
	v.visitStx(g)
}

type PrecGetter struct {
	val Prec
}

func (g *PrecGetter) get(c *Contact) {
	g.val = c.Prec
}

func (g *PrecGetter) accept(v FieldGetterVisitor) {
	v.visitPrec(g)
}

type CkGetter struct {
	val Ck
}

func (g *CkGetter) get(c *Contact) {
	g.val = c.Ck
}

func (g *CkGetter) accept(v FieldGetterVisitor) {
	v.visitCk(g)
}

type SectGetter struct {
	val Sect
}

func (g *SectGetter) get(c *Contact) {
	g.val = c.Sect
}

func (g *SectGetter) accept(v FieldGetterVisitor) {
	v.visitSect(g)
}

type RstRcvdGetter struct {
	val RstRcvd
}

func (g *RstRcvdGetter) get(c *Contact) {
	g.val = c.RstRcvd
}

func (g *RstRcvdGetter) accept(v FieldGetterVisitor) {
	v.visitRstRcvd(g)
}

type RstSentGetter struct {
	val RstSent
}

func (g *RstSentGetter) get(c *Contact) {
	g.val = c.RstSent
}

func (g *RstSentGetter) accept(v FieldGetterVisitor) {
	v.visitRstSent(g)
}

type StateGetter struct {
	val State
}

func (g *StateGetter) get(c *Contact) {
	g.val = c.State
}

func (g *StateGetter) accept(v FieldGetterVisitor) {
	v.visitState(g)
}

type MySotaRefGetter struct {
	val MySotaRef
}

func (g *MySotaRefGetter) get(c *Contact) {
	g.val = c.MySotaRef
}

func (g *MySotaRefGetter) accept(v FieldGetterVisitor) {
	v.visitMySotaRef(g)
}

type MyStateGetter struct {
	val MyState
}

func (g *MyStateGetter) get(c *Contact) {
	g.val = c.MyState
}

func (g *MyStateGetter) accept(v FieldGetterVisitor) {
	v.visitMyState(g)
}

type CntyGetter struct {
	val Cnty
}

func (g *CntyGetter) get(c *Contact) {
	g.val = c.Cnty
}

func (g *CntyGetter) accept(v FieldGetterVisitor) {
	v.visitCnty(g)
}
