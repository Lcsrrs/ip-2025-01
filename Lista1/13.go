package main

import "fmt"

func main() {

	var nota float32

	fmt.Print("Escreva a nota: ")
	fmt.Scan(&nota)

	if nota < 0 {
		fmt.Print("NOTA INVÁLIDA")
	} else if nota < 6 {
		fmt.Print("NOTA: ", nota, "\nCONCEITO: D")
	} else if nota < 7.5 {
		fmt.Print("NOTA: ", nota, "\nCONCEITO: C")
	} else if nota < 9 {
		fmt.Print("NOTA: ", nota, "\nCONCEITO: B")
	} else if nota < 10 {
		fmt.Print("NOTA: ", nota, "\nCONCEITO: A")
	} else {
		fmt.Print("NOTA INVÁLIDA")
	}

}
