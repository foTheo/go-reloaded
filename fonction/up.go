package fonction

import (
	"strings"
)

func Up(mots []string) []string {
	resultat := []string{}
	for _, mot := range mots {
		if mot == "(up)" && len(resultat) > 0 {
			dernierMot := resultat[len(resultat)-1]
			resultat[len(resultat)-1] = strings.ToUpper(dernierMot)
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
