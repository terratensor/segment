package tokenizer

import (
	"testing"

	"github.com/terratensor/segment/rule"
	"github.com/terratensor/segment/segment"
	"github.com/terratensor/segment/split"
)

func TestTokenizer(t *testing.T) {
	splitter := split.NewSplitter(3)
	rules := []rule.Rule{
		rule.NewDashRule(),
		rule.NewFloatRule(),
		rule.NewFractionRule(),
		rule.NewPunctRule(),
		rule.NewOtherRule(),
	}

	tokenizer := NewTokenizer(splitter, rules)

	tests := []struct {
		name     string
		text     string
		expected []segment.Segment
	}{
		{
			name: "Текст с дефисами и числами",
			text: "Кружка-термос на 0.5л (50/64 см³, 516;...)",
			expected: []segment.Segment{
				{Text: "Кружка-термос", Start: 0, End: 13},
				{Text: "на", Start: 14, End: 16},
				{Text: "0.5", Start: 17, End: 20},
				{Text: "л", Start: 20, End: 21},
				{Text: "(", Start: 22, End: 23},
				{Text: "50/64", Start: 23, End: 28},
				{Text: "см³", Start: 29, End: 32},
				{Text: ",", Start: 32, End: 33},
				{Text: "516", Start: 34, End: 37},
				{Text: ";", Start: 37, End: 38},
				{Text: "...", Start: 38, End: 41},
				{Text: ")", Start: 41, End: 42},
			},
		},
		{
			name: "Текст с кириллицей и регистром",
			text: "Пример Текста с Разным Регистром",
			expected: []segment.Segment{
				{Text: "Пример", Start: 0, End: 6},
				{Text: "Текста", Start: 7, End: 13},
				{Text: "с", Start: 14, End: 15},
				{Text: "Разным", Start: 16, End: 22},
				{Text: "Регистром", Start: 23, End: 32},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := tokenizer.Tokenize(tt.text)
			if len(tokens) != len(tt.expected) {
				t.Errorf("Ожидалось %d токенов, получено %d", len(tt.expected), len(tokens))
			}
			for i, token := range tokens {
				if token.Text != tt.expected[i].Text {
					t.Errorf("Токен %d: ожидалось %q, получено %q", i, tt.expected[i].Text, token.Text)
				}
				if token.Start != tt.expected[i].Start {
					t.Errorf("Токен %d: ожидалось начало %d, получено %d", i, tt.expected[i].Start, token.Start)
				}
				if token.End != tt.expected[i].End {
					t.Errorf("Токен %d: ожидалось окончание %d, получено %d", i, tt.expected[i].End, token.End)
				}
			}
		})
	}
}
