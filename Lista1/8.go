package main

import "fmt"

func main() {

	var raio, altura, area, custo float32

	fmt.Print("Escreva o raio e a altura da lata em metros: ")
	fmt.Scan(&raio, &altura)

	area = 2*3.14159*raio*2 + 2*3.14159*raio*altura

	custo = area * 100

	fmt.Print("O VALOR DE CUSTO É: ", custo)

}
