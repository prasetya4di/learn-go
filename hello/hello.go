package main

import (
	"fmt"
	"greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// Get a greeting message and print it.
	message := greetings.Hello("Pras")
	fmt.Println(message)
}
