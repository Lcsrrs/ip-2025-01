package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Escreva dois números: ")
	fmt.Scan(&x, &y)

	if x%2 != 0 {
		fmt.Print("O PRIMEIRO NÚMERO NÃO É PAR")
	} else {
		for i := 0; i < y; i++ {
			fmt.Print(x, "\n")
			x += 2
		}
	}

}
