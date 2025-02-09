package segment

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Константы для символов
const (
	Endings       = ".?!…"
	Dashes        = "‑–—−-"
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
		`(?P<RU>[а-яё]+)|(?P<LAT>[a-z]+)|(?P<INT>\d+)|(?P<PUNCT>[` + regexp.QuoteMeta(Puncts) + `])|(?P<OTHER>\S)`,
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

// Реализация TokenSplit
// 1. Структура TokenSplit

// TokenSplit представляет разбиение текста на левую и правую части.
type TokenSplit struct {
	LeftAtoms  []Atom
	RightAtoms []Atom
	Delimiter  string
}

// NewTokenSplit создаёт новый экземпляр TokenSplit.
func NewTokenSplit(leftAtoms []Atom, delimiter string, rightAtoms []Atom) TokenSplit {
	return TokenSplit{
		LeftAtoms:  leftAtoms,
		RightAtoms: rightAtoms,
		Delimiter:  delimiter,
	}
}

// Left1 возвращает последний атом из LeftAtoms.
func (ts TokenSplit) Left1() Atom {
	if len(ts.LeftAtoms) > 0 {
		return ts.LeftAtoms[len(ts.LeftAtoms)-1]
	}
	return Atom{}
}

// Left2 возвращает предпоследний атом из LeftAtoms.
func (ts TokenSplit) Left2() Atom {
	if len(ts.LeftAtoms) > 1 {
		return ts.LeftAtoms[len(ts.LeftAtoms)-2]
	}
	return Atom{}
}

// Left3 возвращает третий с конца атом из LeftAtoms.
func (ts TokenSplit) Left3() Atom {
	if len(ts.LeftAtoms) > 2 {
		return ts.LeftAtoms[len(ts.LeftAtoms)-3]
	}
	return Atom{}
}

// Right1 возвращает первый атом из RightAtoms.
func (ts TokenSplit) Right1() Atom {
	if len(ts.RightAtoms) > 0 {
		return ts.RightAtoms[0]
	}
	return Atom{}
}

// Right2 возвращает второй атом из RightAtoms.
func (ts TokenSplit) Right2() Atom {
	if len(ts.RightAtoms) > 1 {
		return ts.RightAtoms[1]
	}
	return Atom{}
}

// Right3 возвращает третий атом из RightAtoms.
func (ts TokenSplit) Right3() Atom {
	if len(ts.RightAtoms) > 2 {
		return ts.RightAtoms[2]
	}
	return Atom{}
}
