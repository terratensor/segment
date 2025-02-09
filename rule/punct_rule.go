package rule

import (
	"regexp"
	"strings"

	"github.com/terratensor/segment/segment"
)

// PunctRule реализует правило для знаков препинания.
type PunctRule struct{}

// Apply применяет правило к TokenSplit.
func (r PunctRule) Apply(split segment.TokenSplit) bool {
	if split.Left1().Type != segment.Punct || split.Right1().Type != segment.Punct {
		return false
	}

	left := split.Left1().Text
	right := split.Right1().Text

	// Проверка на смайлики
	if r.isSmile(left + right) {
		return true
	}

	// Проверка на многоточия и комбинации знаков препинания
	if strings.Contains(segment.Endings, left) && strings.Contains(segment.Endings, right) {
		return true
	}

	// Проверка на комбинации типа "--" или "**"
	if left+right == "--" || left+right == "**" {
		return true
	}

	return false
}

// isSmile проверяет, является ли текст смайликом.
func (r PunctRule) isSmile(text string) bool {
	smileRegex := regexp.MustCompile(`^[=:;]-?[)(]{1,3}$`)
	return smileRegex.MatchString(text)
}
