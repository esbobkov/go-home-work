package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"log"
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

	valid, err := isValid(str)
	checkError(err)

	if !valid {
		return "", ErrInvalidString
	}

	var res strings.Builder
	var m int

	s := []rune(str)

	length := len(s)
	for i := 0; i < length; i++ {
		if unicode.IsDigit(s[i]) {
			m, err = strconv.Atoi(string(s[i]))
			checkError(err)

			tmp := strings.Repeat(string(s[i-1]), m)

			_, err = res.WriteString(tmp)
			checkError(err)
		} else if (i == length-1) || (i != length-1 && !unicode.IsDigit(s[i+1])) {
			_, err = res.WriteRune(s[i])
			checkError(err)
		}
	}

	return res.String(), nil
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

func isValid(s string) (bool, error) {
	pattern := `^\d|\d{2,}`
	matched, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		return false, err
	}
	if matched {
		return false, nil
	}

	return true, nil
}
