package main

import "fmt"

func main() {
	var salario float32

	fmt.Print("Escreva o salário: ")
	fmt.Scan(&salario)

	if salario < 300 {
		fmt.Print("VALOR REAJUSTADO: ", salario*1.5)
	} else {
		fmt.Print("VALOR REAJUSTADO: ", salario*1.3)

	}

}
