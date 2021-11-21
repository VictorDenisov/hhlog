package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWritingTemplate(t *testing.T) {
	input := "%f %c %d %t %m"
	r, e := parseWritingTemplate(input)
	assert.Nil(t, e)

	expectedValues := []string{
		"frequency",
		"call",
		"date",
		"time",
		"mode",
	}

	contact := Contact{
		"frequency",
		"call",
		"date",
		"time",
		"mode",
		"srx",
		"stx",
		"prec",
		"ck",
		"sect",
	}

	values := make([]string, len(r))
	vv := &ValueVisitor{}
	for i, s := range r {
		s.get(&contact)
		s.accept(vv)
		values[i] = vv.val
	}

	assert.Equal(t, expectedValues, values)
}

func TestIsTemplateString(t *testing.T) {
	input := "%f\t%err\t%d\t%t"
	assert.True(t, isTemplateString(input))
}

func TestIsTemplateStringCallTime(t *testing.T) {
	input := "%c\t%t"
	assert.True(t, isTemplateString(input))
}

type ValueVisitor struct {
	val string
}

func (v *ValueVisitor) visitFrequency(g *FrequencyGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitCall(g *CallGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitDate(g *DateGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitTime(g *TimeGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitMode(g *ModeGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitBand(g *BandGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitSkcc(g *SkccGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitName(g *NameGetter) {
	v.val = g.val
}

func (v *ValueVisitor) visitSpc(g *SpcGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitSrx(g *SrxGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitStx(g *StxGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitPrec(g *PrecGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitCk(g *CkGetter) {
	v.val = string(g.val)
}

func (v *ValueVisitor) visitSect(g *SectGetter) {
	v.val = string(g.val)
}
