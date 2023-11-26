package main

import (
	"fmt"
	"os"
)

func renderTsv(f *os.File, getters []FieldGetter, contacts []Contact) {
	for _, c := range contacts {
		fmt.Fprintf(f, "V2\t")
		for _, g := range getters {
			fp := NewTsvFieldPrinter()
			g.get(&c)
			g.accept(fp)
			fp.printField(f)
		}
		fmt.Fprintf(f, "\n")
	}
}

type TsvFieldPrinter struct {
	valueVisitor *ValueVisitor
}

func NewTsvFieldPrinter() *TsvFieldPrinter {
	return &TsvFieldPrinter{&ValueVisitor{}}
}

func (fp *TsvFieldPrinter) printField(f *os.File) {
	fmt.Fprintf(f, "%v\t", fp.valueVisitor.val)
}

func (fp *TsvFieldPrinter) visitLiteral(g *LiteralGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitFrequency(g *FrequencyGetter) {
	g.accept(fp.valueVisitor)
	fp.valueVisitor.val = fp.valueVisitor.val + "MHz"
}

func (fp *TsvFieldPrinter) visitCall(g *CallGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitDate(g *DateGetter) {
	g.accept(fp.valueVisitor)
	s := fp.valueVisitor.val
	fp.valueVisitor.val = s[0:4] + "-" + s[4:6] + "-" + s[6:8]
}

func (fp *TsvFieldPrinter) visitTime(g *TimeGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitMode(g *ModeGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitBand(g *BandGetter) {
	g.accept(fp.valueVisitor)
	val := fp.valueVisitor.val
	if val == "20M" {
		fp.valueVisitor.val = "14000"
	}
	if val == "40M" {
		fp.valueVisitor.val = "7000"
	}
}

func (fp *TsvFieldPrinter) visitSkcc(g *SkccGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitName(g *NameGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitSpc(g *SpcGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitSrx(g *SrxGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitStx(g *StxGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitPrec(g *PrecGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitCk(g *CkGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitSect(g *SectGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitRstRcvd(g *RstRcvdGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitRstSent(g *RstSentGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitState(g *StateGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitMySotaRef(g *MySotaRefGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitMyState(g *MyStateGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitCnty(g *CntyGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitMyCall(g *MyCallGetter) {
	g.accept(fp.valueVisitor)
}

func (fp *TsvFieldPrinter) visitMyPotaRef(g *MyPotaRefGetter) {
	g.accept(fp.valueVisitor)
}
