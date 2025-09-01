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
            input:    "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input:    "charmander PIKACHU Bulbasur",
            expected: []string{"charmander", "pikachu", "bulbasur"},
        },
        {
            input:    "no",
            expected: []string{"no"},
        },

        {
            input:    "YES",
            expected: []string{"yes"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        // Check the length of the actual slice against the expected slice
        // if they don't match, use t.Errorf to print an error message
        // and fail the test
        if len(actual) != len(c.expected) {
            t.Errorf("%v != %v", len(actual), len(c.expected))
            t.FailNow()
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            // Check each word in the slice
            // if they don't match, use t.Errorf to print an error message
            // and fail the test
            if word != expectedWord {
                t.Errorf("%v != %v", word, expectedWord)
                t.FailNow()
            }
        }
    }
}