package main

import (
	"fmt"
	"os"
)

func renderAdif(f *os.File, getters []FieldGetter, contacts []Contact) {
	fmt.Fprintf(f, "<adif_ver:5>3.1.2\n")
	fmt.Fprintf(f, "<programid:5>hhlog\n")
	fmt.Fprintf(f, "<programversion:5>0.0.1\n")
	fmt.Fprintf(f, "<EOH>\n")

	for _, c := range contacts {
		for i, g := range getters {
			fp := NewAdifFieldPrinter()
			g.get(&c)
			g.accept(fp)
			if i > 0 {
				fmt.Fprintf(f, "    ")
			}
			fp.printField(f)
		}
		fmt.Fprintf(f, "<EOR>\n")
	}
}

func NewAdifFieldPrinter() *AdifFieldPrinter {
	return &AdifFieldPrinter{valueVisitor: &ValueVisitor{}}
}

type AdifFieldPrinter struct {
	field        string
	valueVisitor *ValueVisitor
}

func (v *AdifFieldPrinter) printField(f *os.File) {
	fmt.Fprintf(f, "<%v:%v>%v\n", v.field, len(v.valueVisitor.val), v.valueVisitor.val)
}

func (v *AdifFieldPrinter) visitFrequency(g *FrequencyGetter) {
	v.field = "FREQ"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitCall(g *CallGetter) {
	v.field = "CALL"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitDate(g *DateGetter) {
	v.field = "QSO_DATE"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitTime(g *TimeGetter) {
	v.field = "TIME_ON"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitMode(g *ModeGetter) {
	v.field = "MODE"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitBand(g *BandGetter) {
	v.field = "BAND"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitSkcc(g *SkccGetter) {
	v.field = "SKCC"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitName(g *NameGetter) {
	v.field = "NAME"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitSpc(g *SpcGetter) {
	v.field = "SPC"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitSrx(g *SrxGetter) {
	v.field = "SRX"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitStx(g *StxGetter) {
	v.field = "STX"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitPrec(g *PrecGetter) {
	v.field = "PRECEDENCE"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitCk(g *CkGetter) {
	v.field = "CHECK"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitSect(g *SectGetter) {
	v.field = "ARRL_SECT"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitRstRcvd(g *RstRcvdGetter) {
	v.field = "RST_RCVD"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitRstSent(g *RstSentGetter) {
	v.field = "RST_SENT"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitState(g *StateGetter) {
	v.field = "STATE"
	g.accept(v.valueVisitor)
}

func (v *AdifFieldPrinter) visitMySotaRef(g *MySotaRefGetter) {
	v.field = "MY_SOTA_REF"
	g.accept(v.valueVisitor)
}
