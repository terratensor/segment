package rule

import (
	"strings"

	"github.com/terratensor/segment/segment"
)

// FloatRule реализует правило для чисел с плавающей точкой.
type FloatRule struct {
	Rule2112
}

// NewFloatRule создаёт новое правило для чисел с плавающей точкой.
func NewFloatRule() Rule {
	return Rule2112{FloatRule{}}
}

// delimiter определяет, является ли текст разделителем для чисел с плавающей точкой.
func (r FloatRule) delimiter(text string) bool {
	return strings.Contains(".,", text)
}

// rule определяет, нужно ли объединять токены.
func (r FloatRule) rule(left, right *segment.Atom) bool {
	if left.Type == segment.Int && right.Type == segment.Int {
		return true
	}
	return false
}
