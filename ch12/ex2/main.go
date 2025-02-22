package main

import (
	"fmt"
	"math/rand"
)

func writeTenNumbersToChannel(ch chan<- int, done chan<- bool) {
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(100)
	}
	done <- true
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan bool)
	go writeTenNumbersToChannel(ch1, done)
	go writeTenNumbersToChannel(ch2, done)

	doneCount := 0
	for {
		select {
		case num := <-ch1:
			fmt.Println("From ch1:", num)
		case num := <-ch2:
			fmt.Println("From ch2:", num)
		case <-done:
			doneCount++
			if doneCount == 2 {
				fmt.Println("All goroutines completed")
				return
			}
		}
	}
}
