package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("Empty name")
	}

	// Return a greeting that embed the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprintf(randomFormat()) // This will fail
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome :)",
		"LOL hello %v",
		"Test, it, out %v",
	}

	return formats[rand.Intn(len(formats))]
}
