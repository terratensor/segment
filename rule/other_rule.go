package rule

import "github.com/terratensor/segment/segment"

// OtherRule реализует правило для других символов.
type OtherRule struct{}

// NewPunctRule создаёт новое правило для других символов.
func NewOtherRule() Rule {
	return OtherRule{}
}

// Apply применяет правило к TokenSplit.
func (r OtherRule) Apply(split segment.TokenSplit) bool {
	left := split.Left1().Type
	right := split.Right1().Type

	if left == segment.Other && (right == segment.Other || right == segment.Ru || right == segment.Lat) {
		return true
	}

	if (left == segment.Other || left == segment.Ru || left == segment.Lat) && right == segment.Other {
		return true
	}

	return false
}
