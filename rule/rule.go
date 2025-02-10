package rule

import "github.com/terratensor/segment/segment"

// Rule представляет интерфейс для правил объединения токенов.
type Rule interface {
	// Apply применяет правило к TokenSplit и возвращает true, если токены нужно объединить.
	Apply(split segment.TokenSplit) bool
}
