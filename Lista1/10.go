package main

import "fmt"

func main() {

	var a, b, c, d, determinante float32

	fmt.Print("Escreva os valores de a, b, c e d da matriz: ")
	fmt.Scan(&a, &b, &c, &d)

	determinante = a*d - b*c

	fmt.Print("O VALOR DO DETERMINANTE É: ", determinante)

}
