package hw02_unpack_string //nolint:golint,stylecheck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	input    string
	expected string
	err      error
}

func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abccd",
			expected: "abccd",
		},
		{
			input:    "3abc",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "aaa0b",
			expected: "aab",
		},
		{
			input:    "a1b2c3",
			expected: "abbccc",
		},
		{
			input:    "诶1必2西3",
			expected: "诶必必西西西",
		},
		{
			input:    "@}2>-5",
			expected: "@}}>-----",
		},
		{
			input:    "зме3д",
			expected: "змееед",
		},
	} {
		t.Run(fmt.Sprintf("%q", tst.input), func(t *testing.T) {
			result, err := Unpack(tst.input)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.expected, result)
		})
	}
}

func TestUnpackWithEscape(t *testing.T) {
	t.Skip() // Remove if task with asterisk completed

	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		result, err := Unpack(tst.input)
		require.Equal(t, tst.err, err)
		require.Equal(t, tst.expected, result)
	}
}
