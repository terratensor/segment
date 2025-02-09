package rule

import "github.com/terratensor/segment"

// Rule представляет интерфейс для правил объединения токенов.
type Rule interface {
	Apply(split segment.TokenSplit) bool
}

// OtherRule реализует правило для объединения токенов типа OTHER.
type OtherRule struct{}

// Apply применяет правило к TokenSplit.
func (r OtherRule) Apply(split segment.TokenSplit) bool {
	left := split.Left1().Type
	right := split.Right1().Type

	if left == segment.Other && (right == segment.Other || right == segment.Ru || right == segment.Lat) {
		return true // JOIN
	}

	if (left == segment.Other || left == segment.Ru || left == segment.Lat) && right == segment.Other {
		return true // JOIN
	}

	return false // SPLIT
}
