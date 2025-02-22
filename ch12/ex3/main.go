package main

import (
	"fmt"
	"math"
	"sync"
)

// buildSqrtMap creates a map of numbers and their square roots
func buildSqrtMap() map[int]float64 {
	result := make(map[int]float64)
	for i := 0; i < 100000; i++ {
		result[i] = math.Sqrt(float64(i))
	}
	return result
}

func main() {
	// Create a cached version of buildSqrtMap using sync.OnceValue
	getSqrtMap := sync.OnceValue(buildSqrtMap)

	// Get the map and look up every 1000th value
	sqrtMap := getSqrtMap()
	for i := 0; i < 100000; i += 1000 {
		fmt.Printf("sqrt(%d) = %.4f\n", i, sqrtMap[i])
	}
}
