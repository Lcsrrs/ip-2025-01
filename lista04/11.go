package main

import (
	f "fmt"
	"math/rand"

	m "math"
)

func main() {
	var vetor []float64
	var soma float64

	for i := 0; i < 100; i++ {
		vetor = append(vetor, rand.Float64())
	}

	for i := 0; i < 50; i++ {
		soma += m.Pow(vetor[i]-vetor[99-i], 3.0)
	}

	f.Print("Soma final: ", soma)
}
