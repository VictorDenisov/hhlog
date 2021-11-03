package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloadSkccRoster(t *testing.T) {
	assert.NotNil(t, DownloadSkccRoster())
}

func TestStripHtmlComments(t *testing.T) {
	value := StripHtmlComments("before<!-- aslkdfj -->after<!-- asldkj -->")
	assert.Equal(t, "beforeafter", value)
}
