package tokenizer

import (
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

	for _, split := range splits {
		if t.shouldJoin(split) {
			buffer += split.Delimiter + split.Right1().Text
		} else {
			if buffer != "" {
				segments = append(segments, segment.Segment{
					Text:  buffer,
					Start: start,
					End:   start + len(buffer),
				})
				buffer = ""
			}
			buffer = split.Right1().Text
			start = split.Right1().Start
		}
	}

	if buffer != "" {
		segments = append(segments, segment.Segment{
			Text:  buffer,
			Start: start,
			End:   start + len(buffer),
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
