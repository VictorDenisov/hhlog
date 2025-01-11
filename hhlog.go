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
		fp := NewHhlogFieldPrinter()
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

func NewHhlogFieldPrinter() *HhlogFieldPrinter {
	return &HhlogFieldPrinter{&ValueVisitor{}}
}

type HhlogFieldPrinter struct {
	valueVisitor *ValueVisitor
}

func (v *HhlogFieldPrinter) printField(f *os.File) {
	fmt.Fprintf(f, "%v\t", v.valueVisitor.val)
}

func (v *HhlogFieldPrinter) visitLiteral(g *LiteralGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitFrequency(g *FrequencyGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitCall(g *CallGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitDate(g *DateGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitTime(g *TimeGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitMode(g *ModeGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitBand(g *BandGetter) {
	g.accept(v.valueVisitor)
	val := v.valueVisitor.val
	if val == "20M" {
		v.valueVisitor.val = "14000"
	}
	if val == "40M" {
		v.valueVisitor.val = "7000"
	}
}

func (v *HhlogFieldPrinter) visitSkcc(g *SkccGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitName(g *NameGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitSpc(g *SpcGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitSrx(g *SrxGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitStx(g *StxGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitPrec(g *PrecGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitCk(g *CkGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitSect(g *SectGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitRstRcvd(g *RstRcvdGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitRstSent(g *RstSentGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitState(g *StateGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitMySotaRef(g *MySotaRefGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitSotaRef(g *SotaRefGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitMyState(g *MyStateGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitCnty(g *CntyGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitMyCall(g *MyCallGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitMyPotaRef(g *MyPotaRefGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitOperator(g *OperatorGetter) {
	g.accept(v.valueVisitor)
}

func (v *HhlogFieldPrinter) visitStationCall(g *StationCallGetter) {
	g.accept(v.valueVisitor)
}

type HhlogFieldNamePrinter struct {
	fieldName string
}

func (v *HhlogFieldNamePrinter) printField(f *os.File) {
	fmt.Fprintf(f, "%v\t", v.fieldName)
}

func (v *HhlogFieldNamePrinter) visitLiteral(g *LiteralGetter) {
	v.fieldName = g.fieldName
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

func (v *HhlogFieldNamePrinter) visitSotaRef(g *SotaRefGetter) {
	v.fieldName = SOTA_REF
}

func (v *HhlogFieldNamePrinter) visitMyState(g *MyStateGetter) {
	v.fieldName = MY_STATE
}

func (v *HhlogFieldNamePrinter) visitCnty(g *CntyGetter) {
	v.fieldName = CNTY
}

func (v *HhlogFieldNamePrinter) visitMyCall(g *MyCallGetter) {
	v.fieldName = MY_CALL
}

func (v *HhlogFieldNamePrinter) visitMyPotaRef(g *MyPotaRefGetter) {
	v.fieldName = MY_POTA_REF
}

func (v *HhlogFieldNamePrinter) visitOperator(g *OperatorGetter) {
	v.fieldName = OPERATOR
}

func (v *HhlogFieldNamePrinter) visitStationCall(g *StationCallGetter) {
	v.fieldName = STATION_CALL
}
