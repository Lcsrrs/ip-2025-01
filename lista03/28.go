package main

import (
	f "fmt"
	m "math"
)

func main() {
	var S float64
	var d = 1

	for i := 0; i < 51; i++ {
		if i%2 == 0 {
			S += 1 / float64((d * d * d))
		} else {
			S -= 1 / float64((d * d * d))
		}
		d += 2
	}

	f.Print("O valor de pi pela fórmula é: ", m.Cbrt(S*32))

}
