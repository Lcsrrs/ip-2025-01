package main

import "fmt"

func main() {
	var base, expoente, potencia int

	fmt.Print("Digite a base: ")
	fmt.Scan(&base)
	potencia = base
	fmt.Print("Digite a potência: ")
	fmt.Scan(&expoente)

	for i := 0; i < expoente-1; i++ {
		potencia *= base
	}

	fmt.Printf("%d^%d = %d", base, expoente, potencia)

}
