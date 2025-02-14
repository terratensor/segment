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
	"github.com/terratensor/segment"
)

func main() {
	tokenizer := segment.NewTokenizer()
	tokens := tokenizer.Tokenize("Кружка-термос на 0.5л (50/64 см³, 516;...)")

	for _, token := range tokens {
		fmt.Printf("Token: %q, Start: %d, End: %d\n", token.Text, token.Start, token.End)
	}
}

```

### Пример использования с указанием подключенных правил
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
### Результат
```bash
go run ./examples/
Token: "Кружка-термос", Start: 0, End: 13
Token: "на", Start: 14, End: 16
Token: "0.5", Start: 17, End: 20
Token: "л", Start: 20, End: 21
Token: "(", Start: 22, End: 23
Token: "50/64", Start: 23, End: 28
Token: "см³", Start: 29, End: 32
Token: ",", Start: 32, End: 33
Token: "516", Start: 34, End: 37
Token: ";", Start: 37, End: 38
Token: "...", Start: 38, End: 41
Token: ")", Start: 41, End: 42
```

## Пакеты

### segment

Пакет `segment` содержит основные структуры и функции для работы с атомами и токенами.

- **Atom**: Минимальная единица текста (слово, число, знак препинания и т.д.).
- **TokenSplit**: Разбиение текста на левую и правую части с разделителем.
- **TokenSplitter**: Разбивает текст на атомы и генерирует `TokenSplit`.

### split

Пакет `split` предоставляет интерфейс для разбиения текста на атомы.

- **Splitter**: Интерфейс для разбиения текста.
- **TokenSplitter**: Реализация интерфейса `Splitter`.

### rule

Пакет `rule` содержит правила для объединения токенов.

- **Rule**: Интерфейс для правил объединения токенов.
- **OtherRule**: Пример правила для объединения токенов типа `OTHER`.

### tokenizer

Пакет `tokenizer` выполняет токенизацию текста с применением правил.

- **Tokenizer**: Основной класс для токенизации текста.

## Лицензия

Этот проект распространяется под лицензией MIT. Подробности см. в файле LICENSE.

## Благодарности

Этот проект вдохновлён библиотекой [razdel](https://github.com/natasha/razdel) на Python, разработанной [Natasha](https://github.com/natasha). Мы выражаем благодарность авторам оригинального проекта за их работу.

---

[![Go Reference](https://pkg.go.dev/badge/github.com/terratensor/segment.svg)](https://pkg.go.dev/github.com/terratensor/segment)