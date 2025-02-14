package tokenizer

import (
	"bufio"
	"os"
	"testing"

	"github.com/terratensor/segment/rule"
	"github.com/terratensor/segment/split"
)

// TestIntegration тестирует токенизацию на данных из файла.
func TestIntegration(t *testing.T) {
	file, err := os.Open("../testdata/tokens.txt")
	if err != nil {
		t.Fatalf("Не удалось открыть файл: %v", err)
	}
	defer file.Close()

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

	tokenizer := NewTokenizer(splitter, rules)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		t.Run(line, func(t *testing.T) {
			tokens := tokenizer.Tokenize(line)

			delimiterCount := 0
			// Подсчитываем количество разделителей | в строке
			for _, token := range tokens {
				if token.Text == "|" {
					delimiterCount++
				}
				t.Logf("Token: %q, Start: %d, End: %d", token.Text, token.Start, token.End)
			}
			// TODO передлать тест,  проччитать строку с разделителями |
			// Потмо убрать разделители ищ строки и еще раз разбить строку, сравнить результаты.
			if len(tokens) != delimiterCount+2 {
				t.Errorf("Tokens: %v", tokens)
			}
			if len(tokens) == 0 {
				t.Errorf("Ожидался хотя бы один токен для строки: %s", line)
			}
		})
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Ошибка при чтении файла: %v", err)
	}
}
