package main

import (
	"fmt"
	"math"
)

func main() {

	var altura, aresta, volume float64

	fmt.Print("Escreva a altura e uma aresta: ")
	fmt.Scan(&altura, &aresta)

	volume = 1.5 * float64(math.Pow(aresta, 2)) * float64(math.Sqrt(3)) * float64(1/3) * altura

	fmt.Print("O VOLUME DA PIRÂMIDE É: ", volume, " METROS CÚBICOS")

}
