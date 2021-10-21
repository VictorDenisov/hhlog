package main

import (
	"bufio"
	_ "errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStructure(t *testing.T) {
	input := "\" %f\t%c\t%d\t%t\t%m\n"
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

func TestReadContacts(t *testing.T) {
	template := "\" %c\t%t\n"
	inputReader := bufio.NewReader(strings.NewReader(template))
	setters, e := readStructure(inputReader)
	assert.Nil(t, e)

	input := "q1bro\t1020\n"
	inputReader = bufio.NewReader(strings.NewReader(input))
	cs, e := readContacts(inputReader, setters)
	assert.Nil(t, e)
	assert.Equal(t, []Contact{Contact{Call: Call("q1bro"), Time: Time("1020")}}, cs)
}

func TestReadContactsWithComment(t *testing.T) {
	template := "\" %c\t%t\n"
	inputReader := bufio.NewReader(strings.NewReader(template))
	setters, e := readStructure(inputReader)
	assert.Nil(t, e)

	input := "q1bro\t1020 \" text in the comment \n"
	inputReader = bufio.NewReader(strings.NewReader(input))
	cs, e := readContacts(inputReader, setters)
	assert.Nil(t, e)
	assert.Equal(t, []Contact{Contact{Call: Call("q1bro"), Time: Time("1020")}}, cs)
}

func TestLineReader(t *testing.T) {
	template := "line\"comment\n"
	lr := NewLineReader(strings.NewReader(template))
	l, c, err := lr.readLine()
	assert.Nil(t, err)
	assert.Equal(t, "comment", c)
	assert.Equal(t, "line", l)
}
