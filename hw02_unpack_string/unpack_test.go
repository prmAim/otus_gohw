package hw02unpackstring

import (
	"errors"
	"testing"

	//nolint:depguard
	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "aaa0b", expected: "aab"},
		{input: "", expected: ""},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `as\\\44`, expected: `as4444`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b", `as\`, `\.`}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestOfNill(t *testing.T) {
	testNils := []struct {
		input       string
		expectedErr error
	}{
		{input: "a4bc2d5e", expectedErr: nil},
		{input: "abccd", expectedErr: nil},
		{input: "3abc", expectedErr: ErrInvalidString},
		{input: "45", expectedErr: ErrInvalidString},
		{input: "aaa10b", expectedErr: ErrInvalidString},
		{input: "aaa0b", expectedErr: nil},
		{input: "", expectedErr: nil},
		{input: "d\n5abc", expectedErr: nil},
		{input: `qwe\4\5`, expectedErr: nil},
		{input: `qwe\45`, expectedErr: nil},
		{input: `qwe\\5`, expectedErr: nil},
		{input: `as\\\44`, expectedErr: nil},
		{input: `as\`, expectedErr: ErrInvalidString},
		{input: `\.`, expectedErr: ErrInvalidString},
	}
	for _, tc := range testNils {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			_, err := Unpack(tc.input)
			if tc.expectedErr == nil {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
			}
		})
	}
}
