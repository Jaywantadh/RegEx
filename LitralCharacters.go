package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	time_start := time.Now()
	fmt.Println("Start time: ", time_start)
	examples := []string{
		"12345", "6789", "01234", "12345678", "7890", "123456789", "34567", "a1b", "a2b", "123-45-6789", "123456789",
	}

	patterns := map[string]string{
		"literal digit '5'":                  "5",
		"literal number '123'":               "123",
		"digit in any position '8'":          "8",
		"literal number '90'":                "90",
		"literal number '456'":               "456",
		"literal digit '2'":                  "2",
		"combining characters 'a1b'":         "a1b",
		"SSN pattern '\\d{3}-\\d{2}-\\d{4}'": `\d{3}-\d{2}-\d{4}`,
	}

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
	time_end := time.Now()

	ex_time := time_end.Sub(time_start)
	fmt.Println("Total execution time:",ex_time)
}
