package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "PIkachu is amazing  pOkEmOn",
			expected: []string{"pikachu", "is", "amazing", "pokemon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("error: \n\tgot: %s \n\texpected: %s", actual, c.expected)
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("not matching!!! \n\t actual: %s \n\t expected: %s", word, expectedWord)
				return
			}
		}
	}
}
