# Segment

Segment — это библиотека для токенизации текста на естественном языке, написанная на Go. Она разбивает текст на атомы (слова, числа, знаки препинания и другие символы) и объединяет их в токены на основе набора правил.

## Особенности

- Поддержка UTF-8 и кириллицы.
- Разбиение текста на атомы с использованием регулярных выражений.
- Гибкая система правил для объединения токенов.
- Простая интеграция в проекты на Go.

## Установка

Для использования библиотеки добавьте её в ваш проект:

```bash
go get github.com/terratensor/segment
```

## Пример использования

```go
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
		rule.NewPunctRule(),
		rule.NewOtherRule(),
		rule.NewYahooRule(),
	}

	tokenizer := tokenizer.NewTokenizer(splitter, rules)
	tokens := tokenizer.Tokenize("Кружка-термос на 0.5л (50/64 см³, 516;...)")

	for _, token := range tokens {
		fmt.Printf("Token: %q, Start: %d, End: %d\n", token.Text, token.Start, token.End)
	}
}
```

## Лицензия

Этот проект распространяется под лицензией MIT. Подробности см. в файле LICENSE.

## Благодарности

Этот проект вдохновлён библиотекой [razdel](https://github.com/natasha/razdel) на Python, разработанной [Natasha](https://github.com/natasha). Мы выражаем благодарность авторам оригинального проекта за их работу.