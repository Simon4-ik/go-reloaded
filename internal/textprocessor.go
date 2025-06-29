package internal

// ProcessText processes the text and applies all transformations.
func ProcessText(input string) string {
	// Hex and bin conversion
	input = HexToDec(input)
	input = BinToDec(input)

	// Case conversion
	input = ConvertCase(input)

	// Format punctuation
	input = FormatPunctuation(input)

	// Fix quotes and articles
	input = FixQuotes(input)
	input = FixArticles(input)

	return input
}
