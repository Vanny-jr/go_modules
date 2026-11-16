package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// SLICE OF NAMES
	names := []string{"Mutua", "Mulwa", "Espana", "Evans", "Gladys"}

	// Get a greeting message and print it.
	// var message string
	messages, err := greetings.Hellos(names)
	// if error occured print in console and then exit the program
	if err != nil {
		log.Fatal(err)
	}
	//if not print the message in console
	fmt.Println(messages)
}
