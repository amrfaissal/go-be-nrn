package gobenrn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBirthDateRFC3339_Success(t *testing.T) {
	dateOfBirth, err := GetBirthDateRFC3339("86081441359")
	assert.Nil(t, err)
	assert.Equal(t, dateOfBirth, "1986-08-14T00:00:00Z")
}

func Test_GetBirthDateRFC3339_Failure(t *testing.T) {
	dateOfBirth, err := GetBirthDateRFC3339("8608144135911")
	assert.Empty(t, dateOfBirth)
	assert.ErrorContains(t, err, ErrInvalidNrnLength.Error())
}
