package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	valid := isValid(str)

	if !valid {
		return "", ErrInvalidString
	}

	var res strings.Builder
	var m int
	var err error

	s := []rune(str)

	length := len(s)
	for i := 0; i < length; i++ {
		if unicode.IsDigit(s[i]) {
			m, err = strconv.Atoi(string(s[i]))
			if err != nil {
				return "", err
			}

			tmp := strings.Repeat(string(s[i-1]), m)

			_, err = res.WriteString(tmp)
			if err != nil {
				return "", err
			}
		} else if (i == length-1) || (i != length-1 && !unicode.IsDigit(s[i+1])) {
			_, err = res.WriteRune(s[i])
			if err != nil {
				return "", err
			}
		}
	}

	return res.String(), nil
}

func isValid(s string) bool {
	re := regexp.MustCompile(`^\d|\d{2,}`)
	return !re.MatchString(s)
}
