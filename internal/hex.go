package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// HexToDec processes all commands recursively until no transformations remain.
func HexToDec(input string) string {
	for {
		// Keep processing commands until the input no longer changes.
		processed := processCommands(input)
		if processed == input {
			break
		}
		input = processed
	}
	return strings.TrimSpace(input)
}

// processCommands applies all transformations (low, up, hex) in a single pass.
func processCommands(input string) string {
	// Handle `(low)` commands: Convert the preceding word to lowercase.
	input = processLowCommand(input)

	// Handle `(up)` commands: Convert the preceding word to uppercase.
	input = processUpCommand(input)

	// Handle `(hex)` commands: Convert hexadecimal numbers to decimal.
	input = processHexCommand(input)

	return input
}

// processLowCommand processes the `(low)` command by converting the preceding word to lowercase.
func processLowCommand(input string) string {
	re := regexp.MustCompile(`([A-Za-z0-9]+)\(\s*low\s*\)`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		parts := re.FindStringSubmatch(match)
		word := parts[1]
		return strings.ToLower(word)
	})
}

// processUpCommand processes the `(up)` command by converting the preceding word to uppercase.
func processUpCommand(input string) string {
	re := regexp.MustCompile(`([A-Za-z0-9]+)\(\s*up\s*\)`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		parts := re.FindStringSubmatch(match)
		word := parts[1]
		return strings.ToUpper(word)
	})
}

// processHexCommand processes the `(hex)` command by converting hexadecimal numbers to decimal.
func processHexCommand(input string) string {
	re := regexp.MustCompile(`([A-Za-z0-9]+)\(\s*hex\s*\)`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		parts := re.FindStringSubmatch(match)
		hexWord := parts[1]
		decimal, err := strconv.ParseInt(hexWord, 16, 64)
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid hex number\n", hexWord)
			return match
		}
		return fmt.Sprintf("%d", decimal)
	})
}
