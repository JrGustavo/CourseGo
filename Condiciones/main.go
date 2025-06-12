package main

import "fmt"

func main() {
	nombre := "Juan"
	edad := 30

	if nombre == "Juan" && edad >= 18 {
		fmt.Println("Hola Juan, eres mayor de edad.")
	} else {
		fmt.Println("Hola, no eres Juan o no eres mayor de edad.")
	}

	if edad >= 18 {
		fmt.Println("Eres mayor de edad.")
	}

	if edad < 18 {
		fmt.Println("Eres menor de edad.")
	} else {
		fmt.Println("Eres adulto.")
	}
}
