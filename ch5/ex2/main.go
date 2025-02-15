package main

import (
	"fmt"
	"log"
	"os"
)

func fileLen(file string) (int, error) {
	content, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer content.Close()

	stat, err := content.Stat()
	fmt.Println(stat.ModTime())
	if err != nil {
		return 0, err
	}
	return int(stat.Size()), nil
}

func main() {
	args := os.Args[1]
	// Get the file length
	length, err := fileLen(args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File length (%s): %d bytes\n", args, length)
}
