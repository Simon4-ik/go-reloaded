package internal

import (
	"regexp"
	"strings"
)

// Function to fix double quotes
func fixDoubleQuotes(input string) string {
	// Remove spaces inside double quotes
	re := regexp.MustCompile(`"\s*(.*?)\s*"`)
	input = re.ReplaceAllString(input, `"$1"`)

	// Remove spaces between quotes and adjacent words
	re = regexp.MustCompile(`(["])\s+([\'\w])`)
	input = re.ReplaceAllString(input, `$1$2`)
	re = regexp.MustCompile(`([\'\w])\s+(["])`)
	input = re.ReplaceAllString(input, `$1$2`)

	// Count the total number of double quotes
	quoteCount := strings.Count(input, `"`)
	var result []rune
	inQuote := false

	// Iterate through the string and fix quotes
	for i := 0; i < len(input); i++ {
		currentChar := rune(input[i])
		if currentChar == '"' {
			// If there is an odd number of quotes and this is the last one, treat it as a closing quote
			if quoteCount%2 != 0 && strings.Count(string(result), `"`) == quoteCount-1 {
				result = append(result, currentChar, ' ') // Add space after closing quote
				continue
			}
			// If this is an opening quote
			if !inQuote {
				inQuote = true
				// Add a space before the opening quote if needed
				if i > 0 && input[i-1] != ' ' && input[i-1] != '"' && input[i-1] != '\'' {
					result = append(result, ' ')
				}
				result = append(result, currentChar)
			} else {
				// If this is a closing quote
				inQuote = false
				result = append(result, currentChar)
				// If there is no space after the closing quote, add a space
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
			}
		} else {
			// Add the current character without changes
			result = append(result, currentChar)
		}
	}
	// Remove unnecessary spaces
	return strings.TrimSpace(string(result))
}

// Function to fix single quotes
func fixSingleQuotes(input string) string {
	// Remove spaces inside single quotes
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	input = re.ReplaceAllString(input, "'$1'")

	// Remove spaces between quotes and adjacent words
	re = regexp.MustCompile(`(['])\s+([\'\w])`)
	input = re.ReplaceAllString(input, `$1$2`)
	re = regexp.MustCompile(`([\'\w])\s+(['])`)
	input = re.ReplaceAllString(input, `$1$2`)

	// List of valid suffixes after single quotes
	validSuffixes := map[string]bool{"t": true, "ll": true, "ve": true, "m": true, "s": true, "d": true, "re": true}

	// Count the total number of single quotes
	quoteCount := strings.Count(input, "'")
	var result []rune
	inQuote := false

	// Iterate through the string and fix quotes
	for i := 0; i < len(input); i++ {
		currentChar := rune(input[i])
		if currentChar == '\'' {
			// Check if the quote is part of a suffix
			if i > 0 && i+1 < len(input) {
				// Find the previous word
				prevWordEnd := i - 1
				for prevWordEnd >= 0 && isLetter(rune(input[prevWordEnd])) {
					prevWordEnd--
				}
				prevWord := input[prevWordEnd+1 : i]
				// Find the next suffix
				nextWordStart := i + 1
				for nextWordStart < len(input) && isLetter(rune(input[nextWordStart])) {
					nextWordStart++
				}
				nextSuffix := input[i+1 : nextWordStart]
				// If the quote is before a valid suffix, skip it as part of the pair
				if isWord(prevWord) && validSuffixes[nextSuffix] {
					result = append(result, currentChar)
					continue
				}
			}

			// If it's the only quote, treat it as a closing one
			if quoteCount == 1 {
				result = append(result, currentChar)
				// Add a space if needed after the quote
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
				continue
			}

			// If there is an odd number of quotes and this is the last one, treat it as closing
			if quoteCount%2 != 0 && strings.Count(string(result), `'`) == quoteCount-1 && !inQuote {
				result = append(result, currentChar)
				// Add a space if needed after the closing quote
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
				continue
			}

			// If this is an opening quote
			if !inQuote {
				inQuote = true
				// Add a space before the opening quote if needed
				if i > 0 && input[i-1] != ' ' && input[i-1] != '\'' && input[i-1] != '"' {
					result = append(result, ' ')
				}
				result = append(result, currentChar)
			} else {
				// If this is a closing quote
				inQuote = false
				result = append(result, currentChar)
				// Add space after closing quote if needed
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
			}
		} else {
			// Add the current character without changes
			result = append(result, currentChar)
		}
	}
	// Remove unnecessary spaces
	return strings.TrimSpace(string(result))
}

// isLetter checks if a character is a letter
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isWord checks if a string is a word (only contains letters)
func isWord(s string) bool {
	for _, ch := range s {
		if !isLetter(ch) {
			return false
		}
	}
	return len(s) > 0
}

// FixQuotes fixes both single and double quotes in the string
func FixQuotes(c string) string {
	c = fixDoubleQuotes(c)
	c = fixSingleQuotes(c)
	return c
}
