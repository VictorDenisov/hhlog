package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	data := `
station:
    call: t6tst
    mail:
        smtp-host: smtp.gmail.com
        port: 587
        password: sample-password
        email: test@gmail.com
pota:
    name: PotaName
    email: PotaMail
wwff:
    name: WwffName
    email: WwffMail
    `
	conf, err := readConfig([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, conf)
}
