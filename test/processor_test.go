package test

import (
	"testing"

	. "go-reloaded/internal"
)

func TestProcessText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello ,world!", "Hello, world!"},
		{"Hello ,world(up)", "Hello, WORLD"},
		{"I was sitting over    !? . there ,and then      BAMM !  !  !", "I was sitting over!?. there, and then BAMM!!!"},
	}

	// Run each test
	for _, tt := range tests {
		result := ProcessText(tt.input)
		if result != tt.expected {
			t.Errorf("Expected '%s', got '%s'", tt.expected, result)
		}
	}
}
