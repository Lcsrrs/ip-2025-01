package main

import f "fmt"

func main() {
	var X, iFat int
	var soma float64

	f.Print("Digite o valor: ")
	f.Scan(&X)

	for i := 0; i < 20; i++ {
		//Determinando i fatorial:
		iFat = 1
		for j := 0; j < i; j++ {
			iFat *= (i - j)
		}

		if i == 0 {
			soma = float64(X)
		} else if i%2 == 0 {
			soma += float64(X) / float64(iFat)
		} else {
			soma -= float64(X) / float64(iFat)
		}

	}

	f.Print("Resultado da soma: ", soma)

}
