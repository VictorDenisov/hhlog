package main

import (
	"fmt"
)

func renderAdif(getters []FieldGetter, contacts []Contact) {
	for _, c := range contacts {
		for i, g := range getters {
			fp := &AdifFieldPrinter{}
			g.get(&c)
			g.accept(fp)
			if i > 0 {
				fmt.Printf("    ")
			}
			fp.printField()
		}
		fmt.Printf("<EOR>\n")
	}
}

type AdifFieldPrinter struct {
	field string
	val   string
}

func (v *AdifFieldPrinter) printField() {
	fmt.Printf("<%v:%v>%v\n", v.field, len(v.val), v.val)
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
