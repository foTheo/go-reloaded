package fonction

import (
	"strings"
	"unicode"
)

func estPonct(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

func Ponctuation(texte string) string {
	runes := []rune(texte)
	out := []rune{}
	i := 0
	for i < len(runes) {
		if !estPonct(runes[i]) {
			out = append(out, runes[i])
			i++
			continue
		}
		avantChiffre := len(out) > 0 && unicode.IsDigit(out[len(out)-1])

		for len(out) > 0 && out[len(out)-1] == ' ' {
			out = out[:len(out)-1]
		}
		debut := i
		for i < len(runes) && estPonct(runes[i]) {
			out = append(out, runes[i])
			i++
		}
		if i < len(runes) && runes[i] != ' ' && runes[i] != '\'' {
			decimal := i-debut == 1 && avantChiffre && unicode.IsDigit(runes[i])
			if !decimal {
				out = append(out, ' ')
			}
		}
	}
	return string(out)
}

func estContraction(runes []rune, i int) bool {
	return i > 0 && i+1 < len(runes) &&
		unicode.IsLetter(runes[i-1]) && unicode.IsLetter(runes[i+1])
}

func Apostrophes(texte string) string {
	runes := []rune(texte)

	total := 0
	for i, r := range runes {
		if r == '\'' && !estContraction(runes, i) {
			total++
		}
	}
	pairs := total - total%2

	out := []rune{}
	vu := 0
	ouvert := false
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r != '\'' || estContraction(runes, i) || vu >= pairs {
			out = append(out, r)
			continue
		}
		vu++
		if !ouvert {
			out = append(out, r)
			for i+1 < len(runes) && runes[i+1] == ' ' {
				i++
			}
			ouvert = true
		} else {
			for len(out) > 0 && out[len(out)-1] == ' ' {
				out = out[:len(out)-1]
			}
			out = append(out, r)
			ouvert = false
		}
	}
	return string(out)
}
