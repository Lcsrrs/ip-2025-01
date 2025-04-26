package main

import (
	f "fmt"
	m "math"
)

func main() {
	var vetor []float64
	var temp float64

	for i := 0; i < 15; i++ {
		f.Printf("Digite o %dº elemento do vetor: ", i+1)
		f.Scan(&temp)
		if temp < 0 {
			vetor = append(vetor, -1)
		} else {
			vetor = append(vetor, m.Sqrt(temp))
		}
	}

	f.Print(vetor)

}
