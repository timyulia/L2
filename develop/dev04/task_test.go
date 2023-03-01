package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack(t *testing.T) {
	var cases = []struct {
		input       []string
		expectedOut map[string][]string
	}{
		{
			input:       []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expectedOut: map[string][]string{"листок": {"слиток", "столик", "листок"}, "пятак": {"пятка", "тяпка", "пятак"}},
		},
		{
			input:       []string{"Пятак", "пЯтка", "Тяпка", "листок"},
			expectedOut: map[string][]string{"пятак": {"пятка", "тяпка", "пятак"}},
		},
		{
			input:       []string{},
			expectedOut: map[string][]string{},
		},
	}

	for _, test := range cases {
		out := devide(test.input)
		assert.Equal(t, test.expectedOut, out)
	}
}
