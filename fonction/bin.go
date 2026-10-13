package fonction

import (
	"strconv"
)

func Bin(mots []string) []string {
	resultat := []string{}
	for _, mot := range mots {
		if mot == "(bin)" && len(resultat) > 0 {
			dernierMot := resultat[len(resultat)-1]
			nombre, err := strconv.ParseInt(dernierMot, 2, 64)
			if err == nil {
				resultat[len(resultat)-1] = strconv.FormatInt(nombre, 10)
			}
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
