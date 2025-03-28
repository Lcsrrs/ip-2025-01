package main

import "fmt"

func main() {
	var (
		x, y, z float32
	)

	fmt.Print("Escreva as notas separadas por espaço: ")
	fmt.Scan(&x, &y, &z)
	var media = (x + y + z) / 3
	fmt.Println("Média das notas: ", media)
	if media >= 6 {
		fmt.Print("APROVADO")
	} else {
		fmt.Print("REPROVADO")
	}
}
