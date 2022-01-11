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
			fp := &AdifFieldPrinter{}
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

type AdifFieldPrinter struct {
	field string
	val   string
}

func (v *AdifFieldPrinter) printField(f *os.File) {
	fmt.Fprintf(f, "<%v:%v>%v\n", v.field, len(v.val), v.val)
}

func (v *AdifFieldPrinter) visitFrequency(g *FrequencyGetter) {
	v.field = "FREQ"
	v.val = g.val
}

func (v *AdifFieldPrinter) visitCall(g *CallGetter) {
	v.field = "CALL"
	v.val = g.val
}

func (v *AdifFieldPrinter) visitDate(g *DateGetter) {
	v.field = "QSO_DATE"
	v.val = g.val
}

func (v *AdifFieldPrinter) visitTime(g *TimeGetter) {
	v.field = "TIME_ON"
	v.val = g.val
}

func (v *AdifFieldPrinter) visitMode(g *ModeGetter) {
	v.field = "MODE"
	v.val = g.val
}

func (v *AdifFieldPrinter) visitBand(g *BandGetter) {
	v.field = "BAND"
	v.val = g.val
}

func (v *AdifFieldPrinter) visitSkcc(g *SkccGetter) {
	v.field = "SKCC"
	v.val = g.val
}

func (v *AdifFieldPrinter) visitName(g *NameGetter) {
	v.field = "NAME"
	v.val = g.val
}

func (v *AdifFieldPrinter) visitSpc(g *SpcGetter) {
	v.field = "SPC"
	v.val = string(g.val)
}

func (v *AdifFieldPrinter) visitSrx(g *SrxGetter) {
	v.field = "SRX"
	v.val = string(g.val)
}

func (v *AdifFieldPrinter) visitStx(g *StxGetter) {
	v.field = "STX"
	v.val = string(g.val)
}

func (v *AdifFieldPrinter) visitPrec(g *PrecGetter) {
	v.field = "PRECEDENCE"
	v.val = string(g.val)
}

func (v *AdifFieldPrinter) visitCk(g *CkGetter) {
	v.field = "CHECK"
	v.val = string(g.val)
}

func (v *AdifFieldPrinter) visitSect(g *SectGetter) {
	v.field = "ARRL_SECT"
	v.val = string(g.val)
}

func (v *AdifFieldPrinter) visitRstRcvd(g *RstRcvdGetter) {
	v.field = "RST_RCVD"
	v.val = string(g.val)
}

func (v *AdifFieldPrinter) visitRstSent(g *RstSentGetter) {
	v.field = "RST_SENT"
	v.val = string(g.val)
}

func (v *AdifFieldPrinter) visitState(g *StateGetter) {
	v.field = "STATE"
	v.val = string(g.val)
}
