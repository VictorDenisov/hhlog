package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeHour_TooBig(t *testing.T) {
	assert.Equal(t, errors.New("Hour should be in the range 00-23"), Time("3346").Valid(), "")
}

func TestTimeHour_Ok(t *testing.T) {
	assert.Equal(t, nil, Time("2346").Valid(), "")
}

func TestTimeMinute_TooBig(t *testing.T) {
	assert.Equal(t, errors.New("Minutes should be in the range 00-59"), Time("2370").Valid(), "")
}
