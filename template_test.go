package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWritingTemplate(t *testing.T) {
	input := "%f\t%c\t%d\t%t\t%m"
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
