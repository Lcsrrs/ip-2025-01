package main

import "fmt"

func main() {
	var x1 int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x1)

	if x1 < 20 || x1 > 90 {
		fmt.Print("O número não está compreendido entre 20 e 90")
	} else {
		fmt.Print("O número está compreendido entre 20 e 90")
	}
}
