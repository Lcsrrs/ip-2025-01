package main

import (
	f "fmt"
	m "math"
)

func main() {
	var x, powX, cosx, iFat float64
	cosx = 1.0

	f.Scan(&x)
	powX = x

	for i := 1; i < 20; i++ {
		//Determinando i fatorial e a potencia de x:
		iFat = 1
		powX = x
		for j := 0; j < 2*i; j++ {
			iFat *= (2*float64(i) - float64(j))
			if j == 0 {
				continue
			} else {
				powX *= x
			}
		}
		if i%2 == 0 {
			cosx += powX / iFat
		} else {
			cosx -= powX / iFat
		}
	}

	f.Printf("%.4f %.4f %.4f", cosx, m.Cos(x), cosx-m.Cos(x))
}
