package main

import (
	"fmt"
)

func renderAdif(getters []FieldGetter, contacts []Contact) {
	for _, c := range contacts {
		for i, g := range getters {
			fp := &FieldPrinter{}
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

type FieldPrinter struct {
	field string
	val   string
}

func (v *FieldPrinter) printField() {
	fmt.Printf("<%v:%v>%v\n", v.field, len(v.val), v.val)
}

func (v *FieldPrinter) visitFrequency(g *FrequencyGetter) {
	v.field = "FREQ"
	v.val = g.val
}

func (v *FieldPrinter) visitCall(g *CallGetter) {
	v.field = "CALL"
	v.val = g.val
}

func (v *FieldPrinter) visitDate(g *DateGetter) {
	v.field = "QSO_DATE"
	v.val = g.val
}

func (v *FieldPrinter) visitTime(g *TimeGetter) {
	v.field = "TIME_ON"
	v.val = g.val
}

func (v *FieldPrinter) visitMode(g *ModeGetter) {
	v.field = "MODE"
	v.val = g.val
}

func (v *FieldPrinter) visitBand(g *BandGetter) {
	v.field = "BAND"
	v.val = g.val
}
