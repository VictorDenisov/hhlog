package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	//data, err := ioutil.ReadFile(".hhlog.conf")
	//if err != nil {
	//	return nil, err
	//}
	data := `
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
