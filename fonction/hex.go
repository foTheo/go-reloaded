package fonction

import (
	"strconv"
)

func Hex(mots []string) []string {
	resultat := []string{}
	for _, mot := range mots {
		if mot == "(hex)" && len(resultat) > 0 {
			dernierMot := resultat[len(resultat)-1]
			nombre, err := strconv.ParseInt(dernierMot, 16, 64)
			if err == nil {
				resultat[len(resultat)-1] = strconv.FormatInt(nombre, 10)
			}
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
