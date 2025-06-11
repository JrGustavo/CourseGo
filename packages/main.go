package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("Home environment variable is not set.")
	} else {
		fmt.Printf("Hola Mundo desde Go! Home: %s\n", envVar)
		os.Exit(0)
	}
}
