package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello     world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "hello, world!",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "hello , world !",
			expected: []string{"hello", ",", "world", "!"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected slice of length %d, got: %d",
				len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("mismatch at index %d: expected %s, got: %s",
					i, expectedWord, word)
			}
		}
	}

}
