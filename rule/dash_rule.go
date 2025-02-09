package rule

import (
	"strings"

	"github.com/terratensor/segment/segment"
)

// DashRule реализует правило для дефисов.
type DashRule struct {
	Rule2112
}

// NewDashRule создаёт новое правило для дефисов.
func NewDashRule() Rule {
	return Rule2112{DashRule{}}
}

// delimiter определяет, является ли текст дефисом.
func (r DashRule) delimiter(text string) bool {
	return strings.Contains(segment.Dashes, text)
}

// rule определяет, нужно ли объединять токены.
func (r DashRule) rule(left, right *segment.Atom) bool {
	if left.Type == segment.Punct || right.Type == segment.Punct {
		return false
	}
	return true
}
