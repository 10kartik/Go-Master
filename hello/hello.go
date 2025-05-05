package main

import (
	"fmt"
	"log"

	"example/hello/greetings" // Local module import
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Using our local greetings package with a single name
	message, err := greetings.Hello("Gopher")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Single greeting:", message)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Using our local greetings package with multiple names
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// Print the map of greetings
	fmt.Println("Multiple greetings:")
	for name, message := range messages {
		fmt.Printf("%v: %v\n", name, message)
	}
}
