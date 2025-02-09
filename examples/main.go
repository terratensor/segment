package main

import (
	"fmt"

	"github.com/terratensor/segment"
)

func main() {
	text := "Кружка-термос на 0.5л (50/64 см³, 516;...)"

	// Создаём TokenSplitter с окном 3
	splitter := segment.NewTokenSplitter(3)

	// Разбиваем текст на TokenSplit
	splits := splitter.Split(text)

	// Выводим результаты
	for _, split := range splits {
		fmt.Printf("Left: %v, Delimiter: %q, Right: %v\n", split.LeftAtoms, split.Delimiter, split.RightAtoms)
	}
}
