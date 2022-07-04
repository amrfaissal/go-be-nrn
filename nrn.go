package gobenrn

import age "github.com/bearbin/go-age"

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

// GetBirthDate returns the date of birth from the NRN (National Registry Number) in the specified time format.
func GetBirthDate(nrn, format string) (string, error) {
	parsedNrn, err := parseNrn(nrn)
	if err != nil {
		return "", err
	}

	birthDate, err := parsedNrn.getBirthDateAsTime()
	if err != nil {
		return "", err
	}

	if format == "" {
		format = simpleDateFormat
	}
	return birthDate.Format(format), nil
}
