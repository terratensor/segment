package rule

import (
	"strings"

	"github.com/terratensor/segment/segment"
)

// FractionRule реализует правило для дробей.
type FractionRule struct {
	Rule2112
}

// NewFractionRule создаёт новое правило для дробей.
func NewFractionRule() Rule {
	return Rule2112{FractionRule{}}
}

// delimiter определяет, является ли текст разделителем для дробей.
func (r FractionRule) delimiter(text string) bool {
	return strings.Contains("/\\", text)
}

// rule определяет, нужно ли объединять токены.
func (r FractionRule) rule(left, right *segment.Atom) bool {
	if left.Type == segment.Int && right.Type == segment.Int {
		return true
	}
	return false
}
