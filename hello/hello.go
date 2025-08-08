package main

import (
    "fmt"
    "log"
    "example.com/greetings"
)
// main function is starting point of program execution
func main() {
    // Get a greeting message and print it.
    log.SetPrefix("greetings:")
    log.SetFlags(0)
    // A slice of names.
    names := []string{"Siwani", "Samantha", "Seetal"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err!=nil {
	log.Fatal(err)
    }
     // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}
