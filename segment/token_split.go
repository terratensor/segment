package segment

// TokenSplit представляет разбиение текста на левую и правую части.
type TokenSplit struct {
	LeftAtoms  []Atom // Атомы слева от разделителя
	RightAtoms []Atom // Атомы справа от разделителя
	Delimiter  string // Разделитель между левой и правой частями
}

// NewTokenSplit создаёт новый экземпляр TokenSplit.
// leftAtoms — атомы слева от разделителя.
// delimiter — разделитель между левой и правой частями.
// rightAtoms — атомы справа от разделителя.
func NewTokenSplit(leftAtoms []Atom, delimiter string, rightAtoms []Atom) TokenSplit {
	return TokenSplit{
		LeftAtoms:  leftAtoms,
		RightAtoms: rightAtoms,
		Delimiter:  delimiter,
	}
}

// Left1 возвращает последний атом из LeftAtoms.
func (ts TokenSplit) Left1() *Atom {
	if len(ts.LeftAtoms) > 0 {
		return &ts.LeftAtoms[len(ts.LeftAtoms)-1]
	}
	return nil
}

// Left2 возвращает предпоследний атом из LeftAtoms.
func (ts TokenSplit) Left2() *Atom {
	if len(ts.LeftAtoms) > 1 {
		return &ts.LeftAtoms[len(ts.LeftAtoms)-2]
	}
	return nil
}

// Left3 возвращает третий с конца атом из LeftAtoms.
func (ts TokenSplit) Left3() *Atom {
	if len(ts.LeftAtoms) > 2 {
		return &ts.LeftAtoms[len(ts.LeftAtoms)-3]
	}
	return nil
}

// Right1 возвращает первый атом из RightAtoms.
func (ts TokenSplit) Right1() *Atom {
	if len(ts.RightAtoms) > 0 {
		return &ts.RightAtoms[0]
	}
	return nil
}

// Right2 возвращает второй атом из RightAtoms.
func (ts TokenSplit) Right2() *Atom {
	if len(ts.RightAtoms) > 1 {
		return &ts.RightAtoms[1]
	}
	return nil
}

// Right3 возвращает третий атом из RightAtoms.
func (ts TokenSplit) Right3() *Atom {
	if len(ts.RightAtoms) > 2 {
		return &ts.RightAtoms[2]
	}
	return nil
}
