package rule

import "github.com/terratensor/segment/segment"

// YahooRule реализует правило для исключения "yahoo!".
type YahooRule struct{}

// Apply применяет правило к TokenSplit.
func (r YahooRule) Apply(split segment.TokenSplit) bool {
	if split.Left1().Normal == "yahoo" && split.Right1().Text == "!" {
		return true
	}
	return false
}
