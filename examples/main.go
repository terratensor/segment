package main

import (
	"fmt"

	"github.com/terratensor/segment"
)

func main() {
	tokenizer := segment.NewTokenizer()
	tokens := tokenizer.Tokenize("Кружка-термос на 0.5л (50/64 см³, 516;...)")

	for _, token := range tokens {
		fmt.Printf("Token: %q, Start: %d, End: %d\n", token.Text, token.Start, token.End)
	}
}
