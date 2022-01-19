package main

import (
	"fmt"
	"strings"
)

func renderCabrillo(getters []FieldGetter, contacts []Contact) {
	fmt.Printf("START-OF-LOG: 3.0\n")
	for _, c := range contacts {
		fmt.Printf("QSO: ")
		for _, g := range getters {
			fp := &CabrilloFieldPrinter{}
			g.get(&c)
			g.accept(fp)
			fp.printField()
		}
		fmt.Printf("\n")
	}
	fmt.Printf("END-OF-LOG:\n")
}

type CabrilloFieldPrinter struct {
	val string
}

func (v *CabrilloFieldPrinter) printField() {
	fmt.Printf("%v\t", v.val)
}

func (v *CabrilloFieldPrinter) visitFrequency(g *FrequencyGetter) {
	f := string(g.val)
	p := strings.Index(f, ".")
	v.val = f[0:p] + f[p+1:p+4]
}

func (v *CabrilloFieldPrinter) visitCall(g *CallGetter) {
	v.val = g.val
}

func (v *CabrilloFieldPrinter) visitDate(g *DateGetter) {
	s := g.val
	v.val = s[0:4] + "-" + s[4:6] + "-" + s[6:8]
}

func (v *CabrilloFieldPrinter) visitTime(g *TimeGetter) {
	v.val = g.val
}

func (v *CabrilloFieldPrinter) visitMode(g *ModeGetter) {
	v.val = g.val
}

func (v *CabrilloFieldPrinter) visitBand(g *BandGetter) {
	if g.val == "20M" {
		v.val = "14000"
	}
	if g.val == "40M" {
		v.val = "7000"
	}
}

func (v *CabrilloFieldPrinter) visitSkcc(g *SkccGetter) {
	v.val = g.val
}

func (v *CabrilloFieldPrinter) visitName(g *NameGetter) {
	v.val = g.val
}

func (v *CabrilloFieldPrinter) visitSpc(g *SpcGetter) {
	v.val = string(g.val)
}

func (v *CabrilloFieldPrinter) visitSrx(g *SrxGetter) {
	v.val = string(g.val)
}

func (v *CabrilloFieldPrinter) visitStx(g *StxGetter) {
	v.val = string(g.val)
}

func (v *CabrilloFieldPrinter) visitPrec(g *PrecGetter) {
	v.val = string(g.val)
}

func (v *CabrilloFieldPrinter) visitCk(g *CkGetter) {
	v.val = string(g.val)
}

func (v *CabrilloFieldPrinter) visitSect(g *SectGetter) {
	v.val = string(g.val)
}

func (v *CabrilloFieldPrinter) visitRstRcvd(g *RstRcvdGetter) {
	v.val = string(g.val)
}

func (v *CabrilloFieldPrinter) visitRstSent(g *RstSentGetter) {
	v.val = string(g.val)
}

func (v *CabrilloFieldPrinter) visitState(g *StateGetter) {
	v.val = string(g.val)
}

func (v *CabrilloFieldPrinter) visitMySotaRef(g *MySotaRefGetter) {
	v.val = string(g.val)
}
