package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/fonction"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		os.Exit(1)
	}
	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	mots := strings.Fields(string(contenu))

	mots = fonction.Bin(mots)
	mots = fonction.Hex(mots)
	mots = fonction.Up(mots)
	mots = fonction.Low(mots)
	mots = fonction.Cap(mots)
	mots = fonction.An(mots)
	mots = fonction.Transformer(mots, "low", strings.ToLower)
	mots = fonction.Transformer(mots, "up", strings.ToUpper)
	mots = fonction.Transformer(mots, "cap", fonction.Capitaliser)

	texte := strings.Join(mots, " ")
	texte = fonction.Ponctuation(texte)
	texte = fonction.Apostrophes(texte)
	if err := os.WriteFile(os.Args[2], []byte(texte), 0644); err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}
}
