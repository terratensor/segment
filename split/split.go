package split

import "github.com/terratensor/segment/segment"

// Splitter представляет интерфейс для разбиения текста на атомы.
type Splitter interface {
	// Split разбивает текст на атомы и возвращает список TokenSplit.
	Split(text string) []segment.TokenSplit
}

// TokenSplitter реализует интерфейс Splitter.
type TokenSplitter struct {
	Window int // Размер окна для анализа атомов
}

// NewSplitter создаёт новый экземпляр TokenSplitter.
// Параметр window определяет размер окна для анализа атомов.
func NewSplitter(window int) *TokenSplitter {
	return &TokenSplitter{
		Window: window,
	}
}

// Split разбивает текст на атомы и возвращает список TokenSplit.
func (ts *TokenSplitter) Split(text string) []segment.TokenSplit {
	return segment.NewTokenSplitter(ts.Window).Split(text)
}
