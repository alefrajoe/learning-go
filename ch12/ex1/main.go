package main

import (
	"log"
	"math/rand"
)

func writeNumberToChannel(number int, channel chan<- int) {
	channel <- number
}

func readNumberFromChannel(channel <-chan int) {
	log.Println("Reading number from channel", <-channel)
}

func channelExercise() {
	channel := make(chan int)
	done := make(chan bool)

	// First writer goroutine
	go func() {
		for j := 0; j < 10; j++ {
			writeNumberToChannel(rand.Intn(100), channel)
		}
		done <- true
	}()

	// Second writer goroutine
	go func() {
		for j := 0; j < 10; j++ {
			writeNumberToChannel(rand.Intn(100), channel)
		}
		done <- true
	}()

	// Third goroutine (reader)
	go func() {
		for i := 0; i < 20; i++ {
			readNumberFromChannel(channel)
		}
		done <- true
	}()

	// Wait for all goroutines to finish
	for i := 0; i < 3; i++ {
		<-done
	}
	close(channel)
}

func main() {
	channelExercise()
	log.Println("Done")
}
