package main

import (
	"bufio"
	_ "errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStructure(t *testing.T) {
	input := "\" %f\t%c\t%d\t%t\t%b\t%m"
	inputReader := bufio.NewReader(strings.NewReader(input))
	r, e := readStructure(inputReader)
	assert.Nil(t, e)
	assert.Equal(t, []string{FREQUENCY, CALL, DATE, TIME, BAND, MODE}, r)
}
