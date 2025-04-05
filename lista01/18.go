package main

import "fmt"

func main() {
	var inicial, razao, elementos int
	var soma = 0

	fmt.Print("Escreva o valor inicial, a razão e o número de elementos da progressão: ")
	fmt.Scan(&inicial, &razao, &elementos)

	for i := 0; i < elementos; i++ {
		soma = soma + inicial
		inicial = +razao
	}

	fmt.Print("Soma:", soma)

}
