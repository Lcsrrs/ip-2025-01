package main

import "fmt"

func main() {
	var soma, elementos int
	var media float64

	for i := 50; i < 71; i += 2 {
		soma += i
		elementos++
	}

	media = float64(soma) / float64(elementos)

	fmt.Printf("Soma = %d\nMédia: %.2f", soma, media)

}
