package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " tesTing Test ",
			expected: []string{"testing", "test"},
		},
		{
			input:    " Test tesTIng ",
			expected: []string{"test", "testing"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("test failed: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("damn girl, thats and error: %v == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
