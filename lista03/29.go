package main

import (
	f "fmt"
)

func main() {
	var n, soma int

	f.Print("Digite o valor de N: ")
	f.Scan(&n)

	for i := 0; i <= n; i++ {
		soma += i
	}

	f.Print("O valor da soma é: ", soma)
}
