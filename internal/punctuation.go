package internal

import (
	"regexp"
	"strings"
)

// Function to format punctuation marks in a string
func FormatPunctuation(input string) string {
	// Remove spaces around punctuation marks
	re := regexp.MustCompile(`\s*([.,!?:;])\s*`)
	input = re.ReplaceAllString(input, "$1")

	// Add a space after punctuation marks if followed by a letter, number, or dash
	re = regexp.MustCompile(`([.,!?:;])([a-zA-Z0-9-])`)
	input = re.ReplaceAllString(input, "$1 $2")

	// Remove double spaces
	input = strings.Join(strings.Fields(input), " ")

	// Remove leading and trailing spaces
	return strings.TrimSpace(input)
}
