package gobenrn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseNrn_Success(t *testing.T) {
	parsedNrn, err := parseNrn(validNrn)
	assert.Nil(t, err)
	assert.Equal(t, parsedNrn, &nrn{Year: "85", Month: "02", Day: "11", Serial: "001", Checksum: "13"})
}

func Test_parseNrn_Failure(t *testing.T) {
	parsedNrn, err := parseNrn(invalidLengthNrn)
	assert.Nil(t, parsedNrn)
	assert.ErrorContains(t, err, ErrInvalidNrnLength.Error())
}
