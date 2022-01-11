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
		"rst_rcvd",
		"rst_sent",
		"state",
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
