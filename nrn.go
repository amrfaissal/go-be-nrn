package gobenrn

import "time"

// GetBirthDateRFC3339 returns the date of birth from the NRN (National Registry Number) in RFC3339 time format.
func GetBirthDateRFC3339(input string) (string, error) {
	parsedNrn, err := parseNrn(input)
	if err != nil {
		return "", err
	}

	birthDate, err := parsedNrn.getBirthDateAsTime()
	if err != nil {
		return "", err
	}
	return birthDate.Format(time.RFC3339), nil
}
