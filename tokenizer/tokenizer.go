package tokenizer

import (
	"unicode/utf8"

	"github.com/terratensor/segment/rule"
	"github.com/terratensor/segment/segment"
	"github.com/terratensor/segment/split"
)

// Tokenizer выполняет токенизацию текста с применением правил.
type Tokenizer struct {
	splitter split.Splitter
	rules    []rule.Rule
}

// NewTokenizer создаёт новый экземпляр Tokenizer.
func NewTokenizer(splitter split.Splitter, rules []rule.Rule) *Tokenizer {
	return &Tokenizer{
		splitter: splitter,
		rules:    rules,
	}
}

// Tokenize выполняет токенизацию текста.
func (t *Tokenizer) Tokenize(text string) []segment.Segment {
	splits := t.splitter.Split(text)
	var segments []segment.Segment
	var buffer string
	start := 0

	// Добавляем первый токен, если он есть
	if len(splits) > 0 {
		firstSplit := splits[0]
		buffer = firstSplit.Left1().Text
		start = firstSplit.Left1().Start
	}

	// Обрабатываем остальные токены
	for _, split := range splits {
		if split.Delimiter == "" && t.shouldJoin(split) {
			// Объединяем токены, добавляя только правый токен
			buffer += split.Right1().Text
		} else {
			if buffer != "" {
				segments = append(segments, segment.Segment{
					Text:  buffer,
					Start: start,
					End:   start + utf8.RuneCountInString(buffer),
				})
				buffer = ""
			}
			buffer = split.Right1().Text
			start = split.Right1().Start
		}
	}

	// Добавляем последний токен, если он есть
	if buffer != "" {
		segments = append(segments, segment.Segment{
			Text:  buffer,
			Start: start,
			End:   start + utf8.RuneCountInString(buffer),
		})
	}

	return segments
}

// shouldJoin определяет, нужно ли объединять токены.
func (t *Tokenizer) shouldJoin(split segment.TokenSplit) bool {
	for _, rule := range t.rules {
		if rule.Apply(split) {
			return true
		}
	}
	return false
}
