package main

import (
	"fmt"

	"github.com/terratensor/segment/rule"
	"github.com/terratensor/segment/split"
	"github.com/terratensor/segment/tokenizer"
)

func main() {
	splitter := split.NewSplitter(3)
	rules := []rule.Rule{
		// rule.DashRule{},
		// rule.FloatRule{},
		// rule.FractionRule{},
		// rule.UnderscoreRule{},
		// rule.PunctRule{},
		rule.OtherRule{},
		// rule.YahooRule{},
	}

	tokenizer := tokenizer.NewTokenizer(splitter, rules)
	tokens := tokenizer.Tokenize("Кружка-термос на 0.5л (50/64 см³, 516;...)")

	for _, token := range tokens {
		fmt.Println(token)
	}
}
