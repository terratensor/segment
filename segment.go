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

// Реализация TokenSplitter
// 1. Структура TokenSplitter

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


// 2. Реализация правил
// Пример реализации правила other

// Rule представляет интерфейс для правил объединения токенов.
type Rule interface {
	Apply(split TokenSplit) bool
}

// OtherRule реализует правило для объединения токенов типа OTHER.
type OtherRule struct{}

// Apply применяет правило к TokenSplit.
func (r OtherRule) Apply(split TokenSplit) bool {
	left := split.Left1().Type
	right := split.Right1().Type

	if left == Other && (right == Other || right == Ru || right == Lat) {
		return true // JOIN
	}

	if (left == Other || left == Ru || left == Lat) && right == Other {
		return true // JOIN
	}

	return false // SPLIT
}

// Реализация Segment и TokenSegmenter

// Segment представляет сегмент текста.
type Segment struct {
	Text  string
	Start int
	End   int
}

// TokenSegmenter выполняет токенизацию текста с применением правил.
type TokenSegmenter struct {
	splitter TokenSplitter
	rules    []Rule
}

// NewTokenSegmenter создаёт новый экземпляр TokenSegmenter.
func NewTokenSegmenter(window int, rules []Rule) TokenSegmenter {
	return TokenSegmenter{
		splitter: NewTokenSplitter(window),
		rules:    rules,
	}
}

// Segment выполняет токенизацию текста.
func (ts TokenSegmenter) Segment(text string) []Segment {
	splits := ts.splitter.Split(text)
	var segments []Segment
	var buffer string
	start := 0

	for _, split := range splits {
		if ts.shouldJoin(split) {
			buffer += split.Delimiter + split.Right1().Text
		} else {
			if buffer != "" {
				segments = append(segments, Segment{
					Text:  buffer,
					Start: start,
					End:   start + len(buffer),
				})
				buffer = ""
			}
			buffer = split.Right1().Text
			start = split.Right1().Start
		}
	}

	if buffer != "" {
		segments = append(segments, Segment{
			Text:  buffer,
			Start: start,
			End:   start + len(buffer),
		})
	}

	return segments
}

// shouldJoin определяет, нужно ли объединять токены.
func (ts TokenSegmenter) shouldJoin(split TokenSplit) bool {
	for _, rule := range ts.rules {
		if rule.Apply(split) {
			return true
		}
	}
	return false
}