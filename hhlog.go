package main

import (
	"fmt"
	"os"
)

func renderHhlog(f *os.File, getters []FieldGetter, contacts []Contact) {
	fmt.Fprintf(f, "\" ")
	fp := &HhlogFieldNamePrinter{}
	for _, g := range getters {
		g.get(&fieldNameContact)
		g.accept(fp)
		fp.printField(f)
	}
	fmt.Fprintf(f, "\n")
	for _, c := range contacts {
		fp := &HhlogFieldPrinter{}
		for _, g := range getters {
			g.get(&c)
			g.accept(fp)
			fp.printField(f)
		}
		fmt.Fprintf(f, "\n")
	}
}

var fieldNameContact = Contact{
	Frequency: FREQUENCY,
	Call:      CALL,
	Date:      DATE,
	Time:      TIME,
	Mode:      MODE,
	Srx:       SRX,
	Stx:       STX,
	Prec:      PREC,
	Ck:        CK,
	Sect:      SECT,
}

type HhlogFieldPrinter struct {
	val string
}

func (v *HhlogFieldPrinter) printField(f *os.File) {
	fmt.Fprintf(f, "%v\t", v.val)
}

func (v *HhlogFieldPrinter) visitFrequency(g *FrequencyGetter) {
	v.val = g.val
}

func (v *HhlogFieldPrinter) visitCall(g *CallGetter) {
	v.val = g.val
}

func (v *HhlogFieldPrinter) visitDate(g *DateGetter) {
	v.val = g.val
}

func (v *HhlogFieldPrinter) visitTime(g *TimeGetter) {
	v.val = g.val
}

func (v *HhlogFieldPrinter) visitMode(g *ModeGetter) {
	v.val = g.val
}

func (v *HhlogFieldPrinter) visitBand(g *BandGetter) {
	if g.val == "20M" {
		v.val = "14000"
	}
	if g.val == "40M" {
		v.val = "7000"
	}
}

func (v *HhlogFieldPrinter) visitSkcc(g *SkccGetter) {
	v.val = g.val
}

func (v *HhlogFieldPrinter) visitName(g *NameGetter) {
	v.val = g.val
}

func (v *HhlogFieldPrinter) visitSpc(g *SpcGetter) {
	v.val = string(g.val)
}

func (v *HhlogFieldPrinter) visitSrx(g *SrxGetter) {
	v.val = string(g.val)
}

func (v *HhlogFieldPrinter) visitStx(g *StxGetter) {
	v.val = string(g.val)
}

func (v *HhlogFieldPrinter) visitPrec(g *PrecGetter) {
	v.val = string(g.val)
}

func (v *HhlogFieldPrinter) visitCk(g *CkGetter) {
	v.val = string(g.val)
}

func (v *HhlogFieldPrinter) visitSect(g *SectGetter) {
	v.val = string(g.val)
}

func (v *HhlogFieldPrinter) visitRstRcvd(g *RstRcvdGetter) {
	v.val = string(g.val)
}

func (v *HhlogFieldPrinter) visitRstSent(g *RstSentGetter) {
	v.val = string(g.val)
}

func (v *HhlogFieldPrinter) visitState(g *StateGetter) {
	v.val = string(g.val)
}

func (v *HhlogFieldPrinter) visitMySotaRef(g *MySotaRefGetter) {
	v.val = string(g.val)
}

type HhlogFieldNamePrinter struct {
	fieldName string
}

func (v *HhlogFieldNamePrinter) printField(f *os.File) {
	fmt.Fprintf(f, "%v\t", v.fieldName)
}

func (v *HhlogFieldNamePrinter) visitFrequency(g *FrequencyGetter) {
	v.fieldName = FREQUENCY
}

func (v *HhlogFieldNamePrinter) visitCall(g *CallGetter) {
	v.fieldName = CALL
}

func (v *HhlogFieldNamePrinter) visitDate(g *DateGetter) {
	v.fieldName = DATE
}

func (v *HhlogFieldNamePrinter) visitTime(g *TimeGetter) {
	v.fieldName = TIME
}

func (v *HhlogFieldNamePrinter) visitMode(g *ModeGetter) {
	v.fieldName = MODE
}

func (v *HhlogFieldNamePrinter) visitBand(g *BandGetter) {
	v.fieldName = BAND
}

func (v *HhlogFieldNamePrinter) visitSkcc(g *SkccGetter) {
	v.fieldName = SKCC
}

func (v *HhlogFieldNamePrinter) visitName(g *NameGetter) {
	v.fieldName = NAME
}

func (v *HhlogFieldNamePrinter) visitSpc(g *SpcGetter) {
	v.fieldName = SPC
}

func (v *HhlogFieldNamePrinter) visitSrx(g *SrxGetter) {
	v.fieldName = SRX
}

func (v *HhlogFieldNamePrinter) visitStx(g *StxGetter) {
	v.fieldName = STX
}

func (v *HhlogFieldNamePrinter) visitPrec(g *PrecGetter) {
	v.fieldName = PREC
}

func (v *HhlogFieldNamePrinter) visitCk(g *CkGetter) {
	v.fieldName = CK
}

func (v *HhlogFieldNamePrinter) visitSect(g *SectGetter) {
	v.fieldName = SECT
}

func (v *HhlogFieldNamePrinter) visitRstRcvd(g *RstRcvdGetter) {
	v.fieldName = RST_RCVD
}

func (v *HhlogFieldNamePrinter) visitRstSent(g *RstSentGetter) {
	v.fieldName = RST_SENT
}

func (v *HhlogFieldNamePrinter) visitState(g *StateGetter) {
	v.fieldName = STATE
}

func (v *HhlogFieldNamePrinter) visitMySotaRef(g *MySotaRefGetter) {
	v.fieldName = MY_SOTA_REF
}
