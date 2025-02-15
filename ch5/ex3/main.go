package main

import "fmt"

func prefixer(prefix string) func(string) string {
	return func(name string) string {
		return prefix + name
	}
}

func main() {
	prefixer := prefixer("Hello, ")
	fmt.Println(prefixer("World!"))
}
