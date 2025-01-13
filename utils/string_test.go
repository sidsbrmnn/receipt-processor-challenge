package utils_test

import (
	"fetch-rewards/receipt-processor-challenge/utils"
	"testing"
)

func TestCountAlphaNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"abc", 3},
		{"123", 3},
		{"abc123", 6},
		{"abc 123", 6},
		{"abc!@#123", 6},
		{"M&M Corner Market", 14},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			res := utils.CountAlphaNumeric(test.input)
			if res != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, res)
			}
		})
	}
}
