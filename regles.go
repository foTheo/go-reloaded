package main

import (
	"strconv"
	"strings"
	"unicode"
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
func Cap(mots []string) []string {
	resultat := []string{}
	for _, mot := range mots {
		if mot == "(cap)" && len(resultat) > 0 {
			lettres := []rune(mot)
			lettres[0] = unicode.ToUpper(lettres[0])
			resultat[len(resultat)-1] = string(lettres)
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
