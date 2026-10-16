package fonction

import "unicode"

func Cap(mots []string) []string {
	resultat := []string{}
	for _, mot := range mots {
		if mot == "(cap)" && len(resultat) > 0 {
			lettres := []rune(resultat[len(resultat)-1])
			lettres[0] = unicode.ToUpper(lettres[0])
			resultat[len(resultat)-1] = string(lettres)
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
