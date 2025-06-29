package internal

import (
	"strings"
	"unicode"
)

func FixArticles(input string) string {
	var result strings.Builder
	lines := strings.Split(input, "\n") // Split text into lines
	for _, line := range lines {
		words := strings.Fields(line) // Split the line into words
		for i := 0; i < len(words)-1; i++ {
			currentWord := words[i]
			nextWord := words[i+1]
			// Check if the next word has more than one character
			if len(nextWord) > 1 {
				// If the article "A" or "a" is before a word starting with a vowel
				if (strings.EqualFold(currentWord, "a") || strings.EqualFold(currentWord, "A")) && isVowel(nextWord[0], nextWord) {
					// Replace with "an" or "AN" based on case
					words[i] = adjustCaseForA(currentWord, nextWord)
				}
				// If the article "AN", "An", "an", or "aN" is before a word starting with a consonant
				if isArticle(currentWord) && !isVowel(nextWord[0], nextWord) {
					// Remove "n" or "N" from the article
					words[i] = removeNFromArticle(currentWord)
				}
			}
		}
		result.WriteString(strings.Join(words, " ") + "\n") // Reassemble the line after all changes
	}
	return strings.TrimSpace(result.String()) // Remove the extra newline at the end
}

// Checks if the word is an article ("a" or "an")
func isArticle(word string) bool {
	return strings.EqualFold(word, "a") || strings.EqualFold(word, "an")
}

var silentHWords = map[string]bool{
	"hour":   true,
	"honest": true,
	"heir":   true,
	"honor":  true,
	"herb":   true,
}

// Checks if the first letter of the word is a vowel or if it's a word with a silent "h"
func isVowel(r byte, nextWord string) bool {
	lowerWord := strings.ToLower(nextWord)
	if lowerWord == "and" || lowerWord == "or" || lowerWord == "an" {
		return false // Do not consider "and" or "an" as starting with a vowel
	}
	// Check if the letter is a vowel or if the word is in the silent "h" list
	if strings.ContainsRune("aeiouAEIOU", rune(r)) || silentHWords[lowerWord] {
		return true
	}
	return false
}

// Adjusts the article "a" to "an" or "AN" depending on the case of the next word
func adjustCaseForA(article string, nextWord string) string {
	if isUpperCase(article) {
		// If the article is uppercase, check if the next word is also uppercase
		if len(nextWord) > 1 && unicode.IsUpper([]rune(nextWord)[1]) {
			return "AN" // If the second letter of the next word is uppercase, use "AN"
		}
		return "An" // If the second letter is lowercase, use "An"
	}
	// In other cases, use "an"
	return "an"
}

// Checks if the first character of the string is uppercase
func isUpperCase(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

// Removes "n" or "N" from the article if the next word starts with a consonant
func removeNFromArticle(article string) string {
	if article == "AN" || article == "An" || article == "aN" || article == "an" {
		return article[:len(article)-1] // Remove "N" or "n"
	}
	return article
}
