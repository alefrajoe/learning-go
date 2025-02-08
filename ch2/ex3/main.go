package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte
	var bigI int
	var bigF float64

	b = math.MaxUint8
	bigI = math.MaxInt64
	bigF = math.MaxFloat64 + 1

	fmt.Println(b, bigI, bigF)
}
