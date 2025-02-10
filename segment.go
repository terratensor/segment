package segment

import (
	"github.com/terratensor/segment/rule"
	"github.com/terratensor/segment/split"
	"github.com/terratensor/segment/tokenizer"
)

// NewTokenizer создаёт новый токенизатор с настройками по умолчанию.
func NewTokenizer() *tokenizer.Tokenizer {
	splitter := split.NewSplitter(3)
	rules := []rule.Rule{
		rule.NewDashRule(),
		rule.NewFloatRule(),
		rule.NewFractionRule(),
		rule.NewUnderscoreRule(),
		rule.NewPunctRule(),
		rule.NewOtherRule(),
		rule.NewYahooRule(),
	}
	return tokenizer.NewTokenizer(splitter, rules)
}
