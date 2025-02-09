package rule

import "github.com/terratensor/segment/segment"

// Rule представляет интерфейс для правил объединения токенов.
type Rule interface {
	Apply(split segment.TokenSplit) bool
}
