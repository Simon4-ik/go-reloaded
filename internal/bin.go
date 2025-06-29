package internal

import (
	"fmt"
	"regexp"
	"strconv"
)

// BinToDec converts binary numbers to decimal.
func BinToDec(input string) string {
	// Regular expression to find binary numbers followed by (bin)
	re := regexp.MustCompile(`\b([01]+)\s?\(bin\)`)

	// Replace all matches with their decimal equivalent
	return re.ReplaceAllStringFunc(input, func(match string) string {
		// Extract the binary value from the match
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match // No binary value found, return original match
		}
		binValue := submatches[1]

		// Convert binary to decimal
		decimalValue, err := strconv.ParseInt(binValue, 2, 64)
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid bin number\n", binValue)
			return match // If conversion fails, return original match
		}

		// Replace the entire match with the decimal value
		return fmt.Sprintf("%d", decimalValue)
	})
}
