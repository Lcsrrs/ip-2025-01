package main

import "fmt"

func main() {
	var x1 int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x1)

	if x1%5 == 0 {
		fmt.Print("O número é divisível por 5")
	} else {
		fmt.Print("O número não é divisível por 5")
	}
}
