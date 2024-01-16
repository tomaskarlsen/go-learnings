package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("Empty name")
	}

	// Return a greeting that embed the name in a message
	message := fmt.Sprintf("Hello, %v. Welcome :)", name)
	return message, nil
}
