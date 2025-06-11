package main

import "fmt"

func main() {
	var numero int = 30
	numero2 := 40
	fmt.Println(numero, numero2)

	var numeroEntero = 10
	var numeroDoble = 20.5
	resultado := numeroEntero + int(numeroDoble)
	fmt.Println(resultado)

	var nombre string = "Ana"
	apellido := "Gomez"
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)

}
