package main

import (
	"fmt"
	"strings"
)

func renderCabrillo(getters []FieldGetter, contacts []Contact) {
	for _, c := range contacts {
		fmt.Printf("QSO: ")
		for _, g := range getters {
			fp := NewCabrilloFieldPrinter()
			g.get(&c)
			g.accept(fp)
			fp.printField()
		}
		fmt.Printf("\n")
	}
}

func NewCabrilloFieldPrinter() *CabrilloFieldPrinter {
	return &CabrilloFieldPrinter{valueVisitor: &ValueVisitor{}}
}

type CabrilloFieldPrinter struct {
	valueVisitor *ValueVisitor
}

func (v *CabrilloFieldPrinter) printField() {
	fmt.Printf("%v\t", v.valueVisitor.val)
}

func (v *CabrilloFieldPrinter) visitLiteral(g *LiteralGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitFrequency(g *FrequencyGetter) {
	g.accept(v.valueVisitor)
	f := v.valueVisitor.val
	p := strings.Index(f, ".")
	v.valueVisitor.val = f[0:p] + f[p+1:p+4]
}

func (v *CabrilloFieldPrinter) visitCall(g *CallGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitDate(g *DateGetter) {
	g.accept(v.valueVisitor)
	s := v.valueVisitor.val
	v.valueVisitor.val = s[0:4] + "-" + s[4:6] + "-" + s[6:8]
}

func (v *CabrilloFieldPrinter) visitTime(g *TimeGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitMode(g *ModeGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitBand(g *BandGetter) {
	g.accept(v.valueVisitor)
	val := v.valueVisitor.val
	if val == "20M" {
		v.valueVisitor.val = "14000"
	}
	if val == "40M" {
		v.valueVisitor.val = "7000"
	}
}

func (v *CabrilloFieldPrinter) visitSkcc(g *SkccGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitName(g *NameGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitSpc(g *SpcGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitSrx(g *SrxGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitStx(g *StxGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitPrec(g *PrecGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitCk(g *CkGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitSect(g *SectGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitRstRcvd(g *RstRcvdGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitRstSent(g *RstSentGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitState(g *StateGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitMySotaRef(g *MySotaRefGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitSotaRef(g *SotaRefGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitMyState(g *MyStateGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitCnty(g *CntyGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitMyCall(g *MyCallGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitMyPotaRef(g *MyPotaRefGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitOperator(g *OperatorGetter) {
	g.accept(v.valueVisitor)
}

func (v *CabrilloFieldPrinter) visitStationCall(g *StationCallGetter) {
	g.accept(v.valueVisitor)
}
