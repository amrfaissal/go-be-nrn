package gobenrn

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	validNrn           = "85021100113"
	validNrnWithSpaces = "850211 001 13"
	validFormattedNrn  = "85.02.11-001.13"
	invalidLengthNrn   = "85021100113019"
)

func Test_GetBirthDateRFC3339_Success(t *testing.T) {
	validNrns := []string{validNrn, validNrnWithSpaces, validFormattedNrn}
	for _, nrn := range validNrns {
		t.Run("With_Valid_NRN_"+nrn, func(t *testing.T) {
			dateOfBirth, err := GetBirthDate(nrn, time.RFC3339)
			assert.Nil(t, err)
			assert.Equal(t, dateOfBirth, "1985-02-11T00:00:00Z")

			dateOfBirth, err = GetBirthDate(nrn, simpleDateFormat)
			assert.Nil(t, err)
			assert.Equal(t, dateOfBirth, "1985-02-11")
		})
	}
}

func Test_GetBirthDateRFC3339_Failure(t *testing.T) {
	dateOfBirth, err := GetBirthDate(invalidLengthNrn, simpleDateFormat)
	assert.Empty(t, dateOfBirth)
	assert.ErrorContains(t, err, ErrInvalidNrnLength.Error())
}

func Test_GetAge_Success(t *testing.T) {
	validNrns := []string{validNrn, validNrnWithSpaces, validFormattedNrn}
	for _, nrn := range validNrns {
		t.Run("With_Valid_NRN_"+nrn, func(t *testing.T) {
			age, err := GetAge(nrn)
			assert.Nil(t, err)
			assert.Equal(t, age, 37)
		})
	}
}

func Test_GetAge_Failure(t *testing.T) {
	age, err := GetAge(invalidLengthNrn)
	assert.Zero(t, age)
	assert.ErrorContains(t, err, ErrInvalidNrnLength.Error())
}
