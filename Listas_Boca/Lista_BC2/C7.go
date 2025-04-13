package main

import (
	f "fmt"
	m "math"
)

func main() {
	var X, iFat int
	var soma float64

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
			soma += m.Pow(float64(X), float64(i)) / float64(iFat)
		} else {
			soma -= m.Pow(float64(X), float64(i)) / float64(iFat)
		}

	}

	f.Print(soma)

}
