package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
)

//go:embed data/*
var data embed.FS

func main() {
	// Define command-line flags
	language := flag.String("l", "italian", "language to use. Options: english, italian, spanish")
	flag.Parse()

	// Load the selected language file
	filePath := fmt.Sprintf("data/%s.txt", *language)

	// Read the file
	file, err := data.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Print the file content
	fmt.Println(string(file))
}
