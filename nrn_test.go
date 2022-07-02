package gobenrn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	validNrn           = "85021100113"
	validNrnWithSpaces = "850211 001 13"
	invalidLengthNrn   = "85021100113019"
)

func Test_GetBirthDateRFC3339_Success(t *testing.T) {
	validNrns := []string{validNrn, validNrnWithSpaces}
	for _, nrn := range validNrns {
		t.Run("With_Valid_NRN_"+nrn, func(t *testing.T) {
			dateOfBirth, err := GetBirthDateRFC3339(nrn)
			assert.Nil(t, err)
			assert.Equal(t, dateOfBirth, "1985-02-11T00:00:00Z")
		})
	}
}

func Test_GetBirthDateRFC3339_Failure(t *testing.T) {
	dateOfBirth, err := GetBirthDateRFC3339(invalidLengthNrn)
	assert.Empty(t, dateOfBirth)
	assert.ErrorContains(t, err, ErrInvalidNrnLength.Error())
}
