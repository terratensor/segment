package rule

import (
	"testing"

	"github.com/terratensor/segment/segment"
)

func TestYahooRule(t *testing.T) {
	rule := YahooRule{}

	tests := []struct {
		name     string
		split    segment.TokenSplit
		expected bool
	}{
		{
			name: "Yahoo и восклицательный знак",
			split: segment.TokenSplit{
				LeftAtoms:  []segment.Atom{{Text: "yahoo", Type: segment.Lat, Normal: "yahoo"}},
				Delimiter:  "",
				RightAtoms: []segment.Atom{{Text: "!", Type: segment.Punct}},
			},
			expected: true,
		},
		{
			name: "Другое слово и восклицательный знак",
			split: segment.TokenSplit{
				LeftAtoms:  []segment.Atom{{Text: "google", Type: segment.Lat, Normal: "google"}},
				Delimiter:  "",
				RightAtoms: []segment.Atom{{Text: "!", Type: segment.Punct}},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rule.Apply(tt.split)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
