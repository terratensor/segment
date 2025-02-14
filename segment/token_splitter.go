package segment

// TokenSplitter разбивает текст на атомы и генерирует TokenSplit.
type TokenSplitter struct {
	Window int // Размер окна для анализа атомов
}

// NewTokenSplitter создаёт новый экземпляр TokenSplitter.
func NewTokenSplitter(window int) TokenSplitter {
	return TokenSplitter{
		Window: window,
	}
}

// Split разбивает текст на атомы и генерирует TokenSplit.
func (ts TokenSplitter) Split(text string) []TokenSplit {
	atoms := Atoms(text)
	var splits []TokenSplit

	if len(atoms) == 1 {
		// Если атомов всего 1, то генерируем TokenSplit
		return append(splits, NewTokenSplit(atoms[0:], "", atoms[0:0]))
	}

	for i := range atoms {
		if i > 0 {
			previous := atoms[i-1]
			delimiter := string([]rune(text)[previous.Stop:atoms[i].Start])
			left := atoms[max(0, i-ts.Window):i]
			right := atoms[i:min(i+ts.Window, len(atoms))]
			splits = append(splits, NewTokenSplit(left, delimiter, right))
		}
	}

	return splits
}

// max возвращает максимальное из двух чисел.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min возвращает минимальное из двух чисел.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
