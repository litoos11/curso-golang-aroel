package main

import (
	"fmt"
	"log"

	"github.com/litoos11/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// message, err := greetings.Hello("Litoos11")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(messages)
}
