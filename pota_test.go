package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInputFile(t *testing.T) {
	call, park, date := parseInputFile("t6tst@k-1123-20211123.hhl")
	assert.Equal(t, Call("t6tst"), call)
	assert.Equal(t, "1123", park)
	assert.Equal(t, Date("20211123"), date)
}
