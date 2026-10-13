package fonction

import (
	"strconv"
	"strings"
)

func N_mots(s string) int {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, ")")
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

func AppliquerN_mots(mots []string, N int, transformation func(string) string) []string {
	debut := len(mots) - N
	if debut < 0 {
		debut = 0
	}
	for i := debut; i < len(mots); i++ {
		mots[i] = transformation(mots[i])
	}
	return mots
}

func Capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(mot)
	return strings.ToUpper(string(lettres[:1])) + strings.ToLower(string(lettres[1:]))
}

func Transformer(mots []string, nomMarqueur string, transformation func(string) string) []string {
	resultat := []string{}
	i := 0

	for i < len(mots) {
		mot := mots[i]

		if mot == "("+nomMarqueur+")" {
			AppliquerN_mots(resultat, 1, transformation)
			i++
		} else if mot == "("+nomMarqueur+"," && i+1 < len(mots) {
			N := N_mots(mots[i+1])
			if N > 0 {
				AppliquerN_mots(resultat, N, transformation)
				i += 2
			} else {
				resultat = append(resultat, mot)
				i++
			}
		} else {
			resultat = append(resultat, mot)
			i++
		}
	}
	return resultat
}
