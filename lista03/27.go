package main

import (
	f "fmt"
	m "math"
)

func main() {
	var x float64
	var iFat int
	cosx := 1.0

	f.Print("Forneça o valor de x: ")
	f.Scan(&x)

	for i := 1; i < 20; i++ {
		//Determinando i fatorial:
		iFat = 1
		for j := 0; j < 2*i; j++ {
			iFat *= (2*i - j)
		}
		if i%2 == 0 {
			cosx += m.Pow(x, 2*float64(i)) / float64(iFat)
		} else {
			cosx -= m.Pow(x, 2*float64(i)) / float64(iFat)
		}
	}

	f.Printf("a) O valor do cosseno de %.2f a partir da fórmula: %.2f\n", x, cosx)
	f.Printf("b) A diferença entre a fórmula e a função Cos(x) = %.2f é de: %.2f", m.Cos(x), cosx-m.Cos(x))
}
