package gobenrn

import (
	"time"

	age "github.com/bearbin/go-age"
)

// GetAge returns person's age from the NRN (National Registry Number).
func GetAge(nrn string) (int, error) {
	parsedNrn, err := parseNrn(nrn)
	if err != nil {
		return 0, err
	}

	birthDate, err := parsedNrn.getBirthDateAsTime()
	if err != nil {
		return 0, err
	}

	return age.Age(*birthDate), nil
}

// GetBirthDateRFC3339 returns the date of birth from the NRN (National Registry Number) in RFC3339 time format.
func GetBirthDateRFC3339(nrn string) (string, error) {
	parsedNrn, err := parseNrn(nrn)
	if err != nil {
		return "", err
	}

	birthDate, err := parsedNrn.getBirthDateAsTime()
	if err != nil {
		return "", err
	}

	return birthDate.Format(time.RFC3339), nil
}
