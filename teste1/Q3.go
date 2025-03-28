package main

import "fmt"

func main() {
	var soma = 0

	for i := 1; i <= 100; i++ {
		fmt.Print(i, "\n")
		soma = soma + i
	}

	fmt.Print("Soma final: ", soma)

}
