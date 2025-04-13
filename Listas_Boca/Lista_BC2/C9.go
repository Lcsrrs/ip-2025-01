package main

import (
	f "fmt"
	m "math"
)

func main() {
	var x float64
	var iFat int
	cosx := 1.0

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

	f.Printf("%.4f %.4f %.4f", cosx, m.Cos(x), cosx-m.Cos(x))
}
