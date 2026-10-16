package fonction

import (
	"strconv"
	"strings"
)

func Nmots(s string) int {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, ")")
	s = strings.TrimSpace(s)
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
	simple := "(" + nomMarqueur + ")"
	debutN := "(" + nomMarqueur + ","

	for i := 0; i < len(mots); i++ {
		mot := mots[i]

		switch {
		case mot == simple:
			AppliquerN_mots(resultat, 1, transformation)

		case mot == debutN && i+1 < len(mots):
			N := Nmots(mots[i+1])
			if N > 0 {
				AppliquerN_mots(resultat, N, transformation)
				i++ // on consomme aussi le token "N)"
			} else {
				resultat = append(resultat, mot)
			}

		case strings.HasPrefix(mot, debutN) && strings.HasSuffix(mot, ")"):
			N := Nmots(strings.TrimPrefix(mot, debutN))
			if N > 0 {
				AppliquerN_mots(resultat, N, transformation)
			} else {
				resultat = append(resultat, mot)
			}

		default:
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
