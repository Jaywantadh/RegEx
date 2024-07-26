// combination.go
// combination.go

// Package main provides a demonstration of using a combination of literal characters and metacharacters in regular expressions.
package main

import (
	"fmt"
	"regexp"
)

// main is the entry point for the program.
// It demonstrates the matching of combinations of literal characters and metacharacters using regular expressions.
func main() {
	examples := []string{
		"apple", "banana", "cat", "cot", "cut", "ct", "cat", "caaaat", 
		"cat", "bat", "rat", "^abc", "abc", "dog", "cat", "bird", "abc", "abcbc", "ac",
		"12345", "6789", "01234", "12345678", "7890", "123456789", "34567",
		"a1b", "a2b", "123-45-6789", "123456789",
	}

	// Map of regular expression patterns to their descriptions
	patterns := map[string]string{
		"literal digit '5'":                  "5",
		"literal number '123'":               "123",
		"digit in any position '8'":          "8",
		"literal number '90'":                "90",
		"literal number '456'":               "456",
		"literal digit '2'":                  "2",
		"combining characters 'a1b'":         "a1b",
		"SSN pattern '\\d{3}-\\d{2}-\\d{4}'": `\d{3}-\d{2}-\d{4}`,
		"dot (.)":                            "c.t",
		"caret (^) at start":                 "^a",
		"dollar ($) at end":                  "t$",
		"asterisk (*) 0 or more 'a's":        "ca*t",
		"plus (+) 1 or more 'a's":            "ca+t",
		"question mark (?) 0 or 1 'a'":       "ca?t",
		"square brackets [cb]":               "[cb]at",
		"negation [^b]":                      "[^b]at",
		"escaped caret (\\^)":                `\^abc`,
		"alternation (|)":                    "cat|dog",
		"grouping (a(bc)+)":                  "a(bc)+",
	}

	// Iterate over each pattern and demonstrate matching against the examples
	for desc, pattern := range patterns {
		fmt.Printf("Pattern: %s (%s)\n", pattern, desc)
		re := regexp.MustCompile(pattern)
		for _, example := range examples {
			if re.MatchString(example) {
				fmt.Printf("  %s -> MATCH\n", example)
			} else {
				fmt.Printf("  %s -> NO MATCH\n", example)
			}
		}
		fmt.Println()
	}
}
