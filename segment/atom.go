package segment

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Константы для символов
const (
	Endings       = ".?!…"
	Dashes        = "‑–—−\\-"
	OpenQuotes    = "«“‘"
	CloseQuotes   = "»”’"
	GenericQuotes = "\"„'"
	Quotes        = OpenQuotes + CloseQuotes + GenericQuotes
	OpenBrackets  = "([{"
	CloseBrackets = ")]}"
	Brackets      = OpenBrackets + CloseBrackets
	Bounds        = Quotes + Brackets
	Puncts        = `\/!#$%&*+,.:;<=>?@^_` + "`|~№…" + Dashes + Quotes + Brackets
)

// Константы для типов атомов
const (
	Ru    = "RU"
	Lat   = "LAT"
	Int   = "INT"
	Punct = "PUNCT"
	Other = "OTHER"
)

// Регулярные выражения
var (
	// ATOM разбивает текст на атомы (слова, числа, знаки препинания и другие символы).
	ATOM = regexp.MustCompile(
		`(?i)(?P<RU>[а-яё]+)|(?P<LAT>[a-z]+)|(?P<INT>\d+)|(?P<PUNCT>[` + regexp.QuoteMeta(Puncts) + `])|(?P<OTHER>\S)`,
	)
)

// Atom представляет минимальную единицу текста.
type Atom struct {
	Start  int
	Stop   int
	Type   string
	Text   string
	Normal string
}

// FindAllRuneSubmatchIndex находит все совпадения регулярного выражения в тексте и возвращает позиции в рунах.
func FindAllRuneSubmatchIndex(re *regexp.Regexp, text string) [][]int {
	var result [][]int
	index := 0

	for i := 0; i < len(text); {
		submatch := re.FindStringIndex(text[index:])
		if submatch == nil {
			break
		}

		// Вычисляем индексы для всех групп
		groups := re.FindStringSubmatchIndex(text[index : index+submatch[1]])
		runeGroups := make([]int, len(groups))
		for j, group := range groups {
			if group == -1 {
				runeGroups[j] = -1 // Группа не найдена
				continue
			}
			// Преобразуем индекс байта в индекс руны
			runeGroups[j] = utf8.RuneCountInString(text[:index+group])
		}

		result = append(result, runeGroups)

		// Перемещаемся к следующей позиции
		i += submatch[1]
		index += submatch[1]
	}

	return result
}

// normalizeText нормализует текст (приводит к нижнему регистру).
func normalizeText(text string) string {
	return strings.ToLower(text)
}

// Atoms генерирует атомы из текста.
func Atoms(text string) []Atom {
	var atoms []Atom
	matches := FindAllRuneSubmatchIndex(ATOM, text)
	for _, match := range matches {
		start, stop := match[0], match[1]
		text := string([]rune(text)[start:stop])

		var type_ string
		for i, name := range ATOM.SubexpNames() {
			if match[i*2] != -1 && name != "" {
				type_ = name
				break
			}
		}

		normal := normalizeText(text)
		atoms = append(atoms, Atom{
			Start:  start,
			Stop:   stop,
			Type:   type_,
			Text:   text,
			Normal: normal,
		})
	}

	return atoms
}
