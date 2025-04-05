package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Escreva um número: ")
	fmt.Scan(&numero)

	if numero < 0 {
		fmt.Print("NEGATIVO\n")
	} else if numero == 0 {
		fmt.Print("NULO\n")

	} else {
		fmt.Print("POSITIVO\n")
	}

}
