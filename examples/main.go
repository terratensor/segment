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
		rule.NewDashRule(),
		rule.NewFloatRule(),
		rule.NewFractionRule(),
		rule.NewUnderscoreRule(),
		rule.PunctRule{},
		rule.OtherRule{},
		rule.YahooRule{},
	}

	tokenizer := tokenizer.NewTokenizer(splitter, rules)
	tokens := tokenizer.Tokenize("Кружка-термос  что-то ружка-термос что-то на Кружка-термос привет мир что-то ни так  0.5л (50/64 см³, 516;...) привет_мир :=)")

	for _, token := range tokens {
		fmt.Println(token)
	}
}
