package error

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// If no name was given, return error
	if name == "" {
		return "", errors.New("No name was given")
	}

	// If name was received return value
	message := fmt.Sprintf("Hello %s", name)
	return message, nil
}
