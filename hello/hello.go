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

    // Request the greeting message
    message, err := greetings.Hello("Gladys")
    if err!=nil {
	log.Fatal(err)
    }
    
    fmt.Println(message)
}
