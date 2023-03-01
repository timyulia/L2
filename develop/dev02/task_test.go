package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack(t *testing.T) {
	var cases = []struct {
		input       string
		expectedOut string
		err         error
	}{
		{
			input:       `a4bc2d5e`,
			expectedOut: `aaaabccddddde`,
			err:         nil,
		},
		{
			input:       `7a4bc2d5e`,
			expectedOut: ``,
			err:         errDigitBegin,
		},
		{
			input:       `av7\v`,
			expectedOut: ``,
			err:         errAfterEscape,
		},
		{
			input:       `qwe\45\`,
			expectedOut: ``,
			err:         errLastEscape,
		},
		{
			input:       `qwe\45`,
			expectedOut: `qwe44444`,
			err:         nil,
		},
		{
			input:       `qwe\\5`,
			expectedOut: `qwe\\\\\`,
			err:         nil,
		},
	}

	for _, test := range cases {
		out, err := unpack(test.input)
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.expectedOut, out)
	}
}
