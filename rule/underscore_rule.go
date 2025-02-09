package rule

import (
	"github.com/terratensor/segment/segment"
)

// UnderscoreRule реализует правило для подчёркиваний.
type UnderscoreRule struct {
	Rule2112
}

// NewUnderscoreRule создаёт новое правило для подчёркиваний.
func NewUnderscoreRule() Rule {
	return Rule2112{UnderscoreRule{}}
}

// delimiter определяет, является ли текст подчёркиванием.
func (r UnderscoreRule) delimiter(text string) bool {
	return text == "_"
}

// rule определяет, нужно ли объединять токены.
func (r UnderscoreRule) rule(left, right *segment.Atom) bool {
	if left.Type == segment.Punct || right.Type == segment.Punct {
		return false
	}
	return true
}
