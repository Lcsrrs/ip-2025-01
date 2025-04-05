package main

import "fmt"

func main() {

	var x1 int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x1)

	if x1%2 == 0 {
		fmt.Print("O número é par")
	} else {
		fmt.Print("O número é ímpar")
	}
}
