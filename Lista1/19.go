package main

import "fmt"

func main() {
	var quantidade int
	var soma float64
	contador := 1.00

	fmt.Print("Escreva um inteiro n: ")
	fmt.Scan(&quantidade)
	for i := 0; i < quantidade; i++ {
		soma = soma + float64(1/contador)
		contador++
		//	fmt.Print("i: ", i, "\nsoma: ", soma, "\ncontador: ", contador, "\n")
	}

	fmt.Print("Soma final: ", soma)

}
