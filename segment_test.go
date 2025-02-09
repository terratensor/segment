package segment

import (
	"testing"
)

func TestTokenSplitter_Split(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		window   int
		expected []TokenSplit
	}{
		{
			name:   "Simple text",
			text:   "Кружка-термос на 0.5л",
			window: 3,
			expected: []TokenSplit{
				{
					LeftAtoms: []Atom{
						{Text: "Кружка", Type: Ru},
						{Text: "-", Type: Punct},
						{Text: "термос", Type: Ru},
					},
					Delimiter: " ",
					RightAtoms: []Atom{
						{Text: "на", Type: Ru},
						{Text: "0.5", Type: Int},
						{Text: "л", Type: Ru},
					},
				},
			},
		},
		{
			name:   "Text with punctuation",
			text:   "Кружка-термос, на 0.5л",
			window: 3,
			expected: []TokenSplit{
				{
					LeftAtoms: []Atom{
						{Text: "Кружка", Type: Ru},
						{Text: "-", Type: Punct},
						{Text: "термос", Type: Ru},
					},
					Delimiter: ",",
					RightAtoms: []Atom{
						{Text: " ", Type: Other},
						{Text: "на", Type: Ru},
						{Text: "0.5", Type: Int},
					},
				},
				{
					LeftAtoms: []Atom{
						{Text: "термос", Type: Ru},
						{Text: ",", Type: Punct},
						{Text: " ", Type: Other},
					},
					Delimiter: " ",
					RightAtoms: []Atom{
						{Text: "на", Type: Ru},
						{Text: "0.5", Type: Int},
						{Text: "л", Type: Ru},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitter := NewTokenSplitter(tt.window)
			splits := splitter.Split(tt.text)

			if len(splits) != len(tt.expected) {
				t.Errorf("Expected %d splits, got %d", len(tt.expected), len(splits))
			}

			for i, split := range splits {
				if len(split.LeftAtoms) != len(tt.expected[i].LeftAtoms) {
					t.Errorf("Split %d: expected %d left atoms, got %d", i, len(tt.expected[i].LeftAtoms), len(split.LeftAtoms))
				}
				if split.Delimiter != tt.expected[i].Delimiter {
					t.Errorf("Split %d: expected delimiter %q, got %q", i, tt.expected[i].Delimiter, split.Delimiter)
				}
				if len(split.RightAtoms) != len(tt.expected[i].RightAtoms) {
					t.Errorf("Split %d: expected %d right atoms, got %d", i, len(tt.expected[i].RightAtoms), len(split.RightAtoms))
				}
			}
		})
	}
}
