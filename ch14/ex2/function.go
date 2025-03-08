package main

import (
	"context"
	"errors"
	"math/rand"
)

func processNumbers(ctx context.Context, result *int, iterations *int) (error) {
	for {
		select {
		case <-ctx.Done():
			return errors.New("context canceled due to timeout")
		default:
			// Generate a random number between 0 and 100_000_000
			current := rand.Intn(100_000_000)
			*result += current
			*iterations++
			if current == 1234 {
				return errors.New("generated 1234")
			}
		}
	}
}
