package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	if a == 0 {
		fmt.Print("Nao e equacao do segundo grau")
		return
	}

	var delta = b*b - 4*a*c
	var x1 float64
	x1 = (-b + math.Sqrt(delta)) / (2 * a)
	var x2 float64
	x2 = (-b - math.Sqrt(delta)) / (2 * a)

	if delta < 0 {
		fmt.Print("Nao ha raizes reais")
	} else {
		fmt.Print(float64(x1), float64(x2))
	}
}
