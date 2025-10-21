package greetings

import (
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, fmt.Errorf("name cannot be empty")
	}
	message := fmt.Sprintf(randomGreeting(), name)
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

func randomGreeting() string {
	greetings := []string{
		"!Hola, %s!",
		"!Bonjour, %s!",
		"!Ciao, %s!",
		"!Hallo, %s!",
	}
	return greetings[rand.Intn(len(greetings))]
}
