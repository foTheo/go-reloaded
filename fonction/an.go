package fonction

import (
	"strings"
	"unicode"
)

func lettre1ere(mot string) bool {
	if mot == "" {
		return false
	}
	lettres := []rune(mot)
	premiere := unicode.ToLower(lettres[0])
	return strings.ContainsRune("aeiouh", premiere)
}

func An(mots []string) []string {
	for i := 0; i < len(mots)-1; i++ {
		if mots[i] == "a" || mots[i] == "A" {
			suivant := mots[i+1]
			if lettre1ere(suivant) {
				if mots[i] == "a" {
					mots[i] = "an"
				} else {
					mots[i] = "An"
				}
			}
		}
	}
	return mots
}
