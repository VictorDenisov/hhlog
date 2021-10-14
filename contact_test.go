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

func TestTime_NonDigital(t *testing.T) {
	assert.Equal(t, errors.New("Time should contain only digits"), Time("ab32").Valid(), "")
}

func TestDate_MonthTooBig(t *testing.T) {
	assert.Equal(t, errors.New("Month should be in the range 1-12"), Date("20213003").Valid(), "")
}

func TestDate_DayTooBig(t *testing.T) {
	assert.Equal(t, errors.New("Day should be in the range 1-31"), Date("20211032").Valid(), "")
}
