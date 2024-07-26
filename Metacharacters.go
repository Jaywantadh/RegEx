// metacharacters.go
////////////////////////////////////////////////////////////////////////////////////////////
// Package main provides a demonstration of using metacharacters in regular expressions.  //
// Imports: 1. fmt base pakage.                                                           //
//          2. regexp pakage for regular expression.                                      //
////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"regexp"
)

// main is the entry point for the program.
// It demonstrates the matching of various metacharacters using regular expressions.
func main() {
	examples := []string{
		"apple", "banana", "cat", "cot", "cut", "ct", "cat", "caaaat", 
		"cat", "bat", "rat", "^abc", "abc", "dog", "cat", "bird", "abc", "abcbc", "ac",
	}

	// Map of regular expression patterns to their descriptions
	patterns := map[string]string{
		"dot (.)":                             "c.t",
		"caret (^) at start":                  "^a",
		"dollar ($) at end":                   "t$",
		"asterisk (*) 0 or more 'a's":         "ca*t",
		"plus (+) 1 or more 'a's":             "ca+t",
		"question mark (?) 0 or 1 'a'":        "ca?t",
		"square brackets [cb]":                "[cb]at",
		"negation [^b]":                       "[^b]at",
		"escaped caret (\\^)":                 `\^abc`,
		"alternation (|)":                     "cat|dog",
		"grouping (a(bc)+)":                   "a(bc)+",
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
