package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
// This is an exported function (starts with capital letter)
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("%s, %v. Welcome!", randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
// This demonstrates returning multiple values and working with maps
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting message formats randomly.
// The returned message format contains a '%v' verb to place the name.
// This is an unexported function (starts with lowercase letter)
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi",
		"Hello",
		"Greetings",
		"Hola",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
