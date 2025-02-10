package rule

import (
	"testing"

	"github.com/terratensor/segment/segment"
)

func TestOtherRule(t *testing.T) {
	rule := OtherRule{}

	tests := []struct {
		name     string
		split    segment.TokenSplit
		expected bool
	}{
		{
			name: "OTHER и RU",
			split: segment.TokenSplit{
				LeftAtoms:  []segment.Atom{{Text: "Δ", Type: segment.Other}},
				Delimiter:  "",
				RightAtoms: []segment.Atom{{Text: "P", Type: segment.Ru}},
			},
			expected: true,
		},
		{
			name: "RU и OTHER",
			split: segment.TokenSplit{
				LeftAtoms:  []segment.Atom{{Text: "mβж", Type: segment.Ru}},
				Delimiter:  "",
				RightAtoms: []segment.Atom{{Text: "Δ", Type: segment.Other}},
			},
			expected: true,
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
