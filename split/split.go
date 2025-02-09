package split

import "github.com/terratensor/segment/segment"

// Splitter представляет интерфейс для разбиения текста на атомы.
type Splitter interface {
	Split(text string) []segment.TokenSplit
}

// TokenSplitter реализует интерфейс Splitter.
type TokenSplitter struct {
	Window int
}

// NewSplitter создаёт новый экземпляр TokenSplitter.
func NewSplitter(window int) *TokenSplitter {
	return &TokenSplitter{
		Window: window,
	}
}

// Split разбивает текст на атомы и генерирует TokenSplit.
func (ts *TokenSplitter) Split(text string) []segment.TokenSplit {
	return segment.NewTokenSplitter(ts.Window).Split(text)
}
