package main

import (
	"fmt"
	"os"
	"strings"
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
	fichier, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer fichier.Close()

	mots := strings.Fields(string(contenu))

	fmt.Println(mots)
}
