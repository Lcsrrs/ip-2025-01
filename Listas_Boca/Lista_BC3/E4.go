package main

import (
	f "fmt"
	m "math"
)

func main() {
	var vetor []float64
	var soma, temp float64
	var quantidade int

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		f.Scan(&temp)
		vetor = append(vetor, temp)
	}

	for i := 0; i < quantidade/2; i++ {
		soma += m.Pow(vetor[i]-vetor[quantidade-i-1], 3.0)
	}

	f.Print(soma)
}
