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
	for i, s := range r {
		values[i] = s(&contact)
	}

	assert.Equal(t, expectedValues, values)
}
