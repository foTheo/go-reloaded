package fonction

import (
	"strings"
)

func Low(mots []string) []string {
	resultat := []string{}
	for _, mot := range mots {
		if mot == "(low)" && len(resultat) > 0 {
			dernierMot := resultat[len(resultat)-1]
			resultat[len(resultat)-1] = strings.ToLower(dernierMot)
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
