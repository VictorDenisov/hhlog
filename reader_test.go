package main

import (
	"bufio"
	_ "errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStructure(t *testing.T) {
	input := "\" %f\t%c\t%d\t%t\t%m"
	inputReader := bufio.NewReader(strings.NewReader(input))
	r, e := readStructure(inputReader)
	assert.Nil(t, e)

	expectedContacts := []Contact{
		Contact{Frequency: Frequency("test")},
		Contact{Call: Call("test")},
		Contact{Date: Date("test")},
		Contact{Time: Time("test")},
		Contact{Mode: Mode("test")},
	}

	contacts := make([]Contact, len(r))
	for i, s := range r {
		s(&contacts[i], "test")
	}

	assert.Equal(t, expectedContacts, contacts)
}
