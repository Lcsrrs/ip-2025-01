package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Print("Digite os coeficiente A, B e C da função: ")
	fmt.Scan(&a, &b, &c)

	var delta = b*b - 4*a*c
	var x1 = (-b + math.Sqrt(delta)) / (2 * a)
	var x2 = (-b - math.Sqrt(delta)) / (2 * a)

	if delta < 0 {
		fmt.Print("Raízes imaginárias")
	} else if delta == 0 {
		fmt.Print("Raíz única e igual a ", x1)
	} else {
		fmt.Printf("Raízes reais e distintas e iguais a %f e %f", x1, x2)
	}
}
