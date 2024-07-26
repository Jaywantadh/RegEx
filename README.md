# Regular Expression and Patternmatching Computations with Golang

This repository contains Go programs that demonstrate the usage of regular expressions with literal characters, metacharacters, and their combinations. Each program showcases how different patterns match or do not match various test strings.

## Files

1. `LiteralCharacters.go`
2. `Metacharacters.go`
3. `Combination.go`

### `literal_characters.go`

This program demonstrates the use of literal characters in regular expressions. Literal characters match exactly what you type.

#### Patterns Used

- `"5"`: Matches the literal digit '5'.
- `"123"`: Matches the literal number '123'.
- `"8"`: Matches the digit '8' in any position.
- `"90"`: Matches the literal number '90'.
- `"456"`: Matches the literal number '456'.
- `"2"`: Matches the literal digit '2'.
- `"a1b"`: Matches the string 'a1b'.
- `\d{3}-\d{2}-\d{4}`: Matches a Social Security Number (SSN) pattern.

### `metacharacters.go`

This program demonstrates the use of metacharacters in regular expressions. Metacharacters have special meanings and are used to create complex patterns.

#### Patterns Used

- `"c.t"`: Matches any single character except a newline between 'c' and 't'.
- `"^a"`: Matches if the string starts with 'a'.
- `"t$"`: Matches if the string ends with 't'.
- `"ca*t"`: Matches 'c' followed by 0 or more 'a's and then 't'.
- `"ca+t"`: Matches 'c' followed by 1 or more 'a's and then 't'.
- `"ca?t"`: Matches 'c' followed by 0 or 1 'a' and then 't'.
- `"[cb]at"`: Matches 'cat' or 'bat'.
- `"[^b]at"`: Matches any string ending with 'at' except starting with 'b'.
- `\^abc`: Matches the literal string '^abc'.
- `"cat|dog"`: Matches 'cat' or 'dog'.
- `"a(bc)+"`: Matches 'a' followed by one or more 'bc'.

### `combination.go`

This program demonstrates the use of a combination of literal characters and metacharacters in regular expressions.

#### Patterns Used

- All patterns from `literal_characters.go` and `metacharacters.go`.

## Running the Programs

To run the programs, use the following commands:

```sh
go run LiteralCharacters.go
go run Metacharacters.go
go run Combination.go
