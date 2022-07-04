package gobenrn

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	ErrInvalidNrnLength  = errors.New("invalid_nrn_length")
	ErrInvalidBirthYear  = errors.New("invalid_birth_year")
	ErrInvalidBirthMonth = errors.New("invalid_birth_month")
	ErrInvalidBirthDay   = errors.New("invalid_birth_day")
	ErrUnknownBirthDate  = errors.New("unknown_birth_date")
)

const (
	validNrnLength                 = 11
	simpleDateFormat               = "2006-01-02"
	bisMonthIncrementGenderUnknown = 20
)

type nrn struct {
	Year     string
	Month    string
	Day      string
	Serial   string
	Checksum string
}

func (input *nrn) getBirthYear() (int, error) {
	var year int

	checksum19, err := mod97(input.String() + input.Serial)
	if err != nil {
		return 0, err
	}
	checksum20, err := mod97("2" + input.String() + input.Serial)
	if err != nil {
		return 0, err
	}

	switch {
	case checksum19 == input.Checksum:
		year, _ = strconv.Atoi("19" + input.Year)
	case checksum20 == input.Checksum:
		year, _ = strconv.Atoi("20" + input.Year)
	}
	return year, nil
}

func (input *nrn) getBirthMonth() (int, error) {
	birthMonth, err := strconv.Atoi(input.Month)
	if err != nil {
		return 0, err
	}

	for birthMonth >= bisMonthIncrementGenderUnknown {
		birthMonth -= bisMonthIncrementGenderUnknown
	}
	return birthMonth, nil
}

func (input *nrn) getBirthDay() (int, error) {
	birthDay, err := strconv.Atoi(input.Day)
	if err != nil {
		return 0, err
	}
	return birthDay, nil
}

func (input *nrn) getBirthDateAsTime() (*time.Time, error) {
	year, err := input.getBirthYear()
	if err != nil {
		return nil, ErrInvalidBirthYear
	}

	month, err := input.getBirthMonth()
	if err != nil {
		return nil, ErrInvalidBirthMonth
	}

	day, err := input.getBirthDay()
	if err != nil {
		return nil, ErrInvalidBirthDay
	}

	if year == 0 || month == 0 || day == 0 {
		return nil, ErrUnknownBirthDate
	}

	dateOfBirthSimple := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return &dateOfBirthSimple, nil
}

func (input *nrn) String() string {
	return input.Year + input.Month + input.Day
}

func parseNrn(input string) (*nrn, error) {
	regex, err := regexp.Compile(`[^\d]+`)
	if err != nil {
		return nil, err
	}

	normalizedNrn := regex.ReplaceAllString(input, "")

	if len(normalizedNrn) != validNrnLength {
		return nil, ErrInvalidNrnLength
	}

	birthDate := normalizedNrn[0:6]
	return &nrn{
		Year:     birthDate[0:2],
		Month:    birthDate[2:4],
		Day:      birthDate[4:],
		Serial:   normalizedNrn[6:9],
		Checksum: normalizedNrn[9:11],
	}, nil
}

func mod97(s string) (string, error) {
	inputAsInt, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}

	mod := 97 - (inputAsInt % 97)
	return fmt.Sprintf("%02d", mod), nil
}
