package rule

import "github.com/terratensor/segment/segment"

// Rule2112 представляет базовое правило для работы с разделителями.
type Rule2112 struct{}

// Apply применяет правило к TokenSplit.
func (r Rule2112) Apply(split segment.TokenSplit) bool {
	left, right := r.getLeftRight(split)
	if left == nil || right == nil {
		return false
	}
	return r.rule(left, right)
}

// getLeftRight возвращает левый и правый атомы для анализа.
func (r Rule2112) getLeftRight(split segment.TokenSplit) (*segment.Atom, *segment.Atom) {
	if r.delimiter(split.Left1().Text) {
		// Пример: cho-|to
		return split.Left2(), split.Right1()
	} else if r.delimiter(split.Right1().Text) {
		// Пример: cho|-to
		return split.Left1(), split.Right2()
	}
	return nil, nil
}

// delimiter определяет, является ли текст разделителем.
func (r Rule2112) delimiter(text string) bool {
	// Этот метод должен быть переопределён в дочерних правилах.
	return false
}

// rule определяет, нужно ли объединять токены.
func (r Rule2112) rule(left, right *segment.Atom) bool {
	// Этот метод должен быть переопределён в дочерних правилах.
	return false
}
