package gobenrn

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	validNrn                = "85021100113"
	validNrnWithSpaces      = "850211 001 13"
	validFormattedNrn       = "85.02.11-001.13"
	validFormattedFemaleNrn = "86.02.15-002.10"
	invalidLengthNrn        = "85021100113019"
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
			assert.NotZero(t, age)
		})
	}
}

func Test_GetAge_Failure(t *testing.T) {
	age, err := GetAge(invalidLengthNrn)
	assert.Zero(t, age)
	assert.ErrorContains(t, err, ErrInvalidNrnLength.Error())
}

func Test_IsMale_Success(t *testing.T) {
	validNrns := []string{validNrn, validNrnWithSpaces, validFormattedNrn}
	for _, nrn := range validNrns {
		t.Run("With_Valid_NRN_"+nrn, func(t *testing.T) {
			male, err := IsMale(nrn)
			assert.Nil(t, err)
			assert.True(t, male)
		})
	}
}

func Test_IsMale_Failure(t *testing.T) {
	input := strings.Replace(validFormattedNrn, "001", "002", 1)
	male, err := IsMale(input)
	assert.Nil(t, err)
	assert.False(t, male)
}

func Test_IsFemale_Success(t *testing.T) {
	female, err := IsFemale(validFormattedFemaleNrn)
	assert.Nil(t, err)
	assert.True(t, female)
}

func Test_IsFemale_Failure(t *testing.T) {
	input := strings.Replace(validFormattedFemaleNrn, "002", "003", 1) // Females have an even serial number
	female, err := IsFemale(input)
	assert.Nil(t, err)
	assert.False(t, female)
}
