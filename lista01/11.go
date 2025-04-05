package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Escreva um número: ")
	fmt.Scan(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Print("O NÚMERO É DIVISÍVEL")
	} else {
		fmt.Print("O NÚMERO NÃO É DIVISÍVEL")
	}

}
