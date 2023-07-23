package gobenrn

import (
	"reflect"
	"strconv"

	age "github.com/bearbin/go-age"
)

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

// IsMale checks if the NRN holder is a Male.
func IsMale(nrn string) (bool, error) {
	parsedNrn, err := parseNrn(nrn)
	if err != nil {
		return false, err
	}

	serial, _ := strconv.Atoi(parsedNrn.Serial)
	return serial%2 != 0, nil
}

// IsFemale checks if the NRN holder is a Female.
func IsFemale(nrn string) (bool, error) {
	parsedNrn, err := parseNrn(nrn)
	if err != nil {
		return false, err
	}

	serial, _ := strconv.Atoi(parsedNrn.Serial)
	return serial%2 == 0, nil
}

// IsBirthDateKnown checks if the NRN holder's birth date is known.
func IsBirthDateKnown(nrn string) (bool, error) {
	parsedNrn, err := parseNrn(nrn)
	if err != nil {
		return false, err
	}

	month, _ := parsedNrn.getBirthMonth()
	day, _ := parsedNrn.getBirthDay()
	return month != 0 && day != 0, nil
}

// Equal checks if two differently formatted NRNs are the same.
func Equal(nrn1, nrn2 string) (bool, error) {
	parsedNrn1, err := parseNrn(nrn1)
	if err != nil {
		return false, err
	}

	parsedNrn2, err := parseNrn(nrn2)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(parsedNrn1, parsedNrn2), nil
}
