package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result := 0
	iterations := 0

	err := processNumbers(ctx, &result, &iterations)
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("Result:", result)
	log.Println("Iterations:", iterations)
}
