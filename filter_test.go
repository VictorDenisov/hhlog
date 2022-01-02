package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExpr(t *testing.T) {
	expr, err := ParseExpr("%c == q6qtx")
	assert.Nil(t, err)
	assert.Equal(t, "q6qtx", expr.(*Eq).value)
	r := expr.run(&Contact{
		Call: "q6qtx",
	})
	assert.True(t, r)
}

func TestParseExprLess(t *testing.T) {
	expr, err := ParseExpr("%d <= 20211220")
	assert.Nil(t, err)
	r := expr.run(&Contact{
		Date: "20211219",
	})
	assert.True(t, r)
}

func TestParseExprMore(t *testing.T) {
	expr, err := ParseExpr("%d >= 20211220")
	assert.Nil(t, err)
	r := expr.run(&Contact{
		Date: "20211221",
	})
	assert.True(t, r)
}

func TestParseExprNotEqual(t *testing.T) {
	expr, err := ParseExpr("%c == q6qtx")
	assert.Nil(t, err)
	assert.Equal(t, "q6qtx", expr.(*Eq).value)
	r := expr.run(&Contact{
		Call: "q6qtt",
	})
	assert.False(t, r)
}

func TestParseFilter(t *testing.T) {
	filter, err := ParseFilter("%f == 14.112 && %c == q6qtx")
	assert.Nil(t, err)
	r := filter.run(&Contact{
		Frequency: "14.112",
		Call:      "q6qtx",
	})
	assert.True(t, r)
}
