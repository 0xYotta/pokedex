package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Printf("Case: %s\n", c.input)
		fmt.Printf("Expecting: %s\n", c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("error: wrong length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			fmt.Printf("word: %s || expectedWord: %s", word, expectedWord)
			if word != expectedWord {
				fmt.Print("\tFAIL!\n")
				t.Errorf("\nerror: words not matching! word: %s != expected: %s", word, expectedWord)
			}
			fmt.Print("\tOK!\n")
		}
	}
}
