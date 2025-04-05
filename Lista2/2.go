package main

import "fmt"

func main() {

	var x1 int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x1)

	if x1 < 0 {
		fmt.Print("O número é negativo")
	} else if x1 == 0 {
		fmt.Print("O número é nulo")
	} else {
		fmt.Print("O número é positivo")
	}

}
