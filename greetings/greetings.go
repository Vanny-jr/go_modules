package greetings

import (
	"errors" // Deals with errors
	"fmt"    // Deals with input and output formatting

	"math/rand" // Generate random numbers
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}


// Hellos function is added and it has a slice of names ratheer than single names.
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
	// A slice of message formats. Omiting the size of the slice while declaring it
	// like this : []string{} tells Go that the size of the array underlying the  slice can dynammicaly change
	formats := []string{
		"Hola, %v. Bienvenido!",
		"Great to see you, %v!",
		"Hail, %v! Well meeting you!",
	}

	// Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
