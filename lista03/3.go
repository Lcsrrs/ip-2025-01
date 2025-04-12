package main

import "fmt"

func main() {
	var salCarlos, salJoao float64
	var meses int

	fmt.Print("Digite o salário de Carlos: ")
	fmt.Scan(&salCarlos)
	salJoao = salCarlos / 3

	for salCarlos >= salJoao {
		salCarlos *= 1.02
		salJoao *= 1.05
		meses++
	}

	fmt.Print("A quantidade de meses para o salário de João superar o de Carlos é: ", meses)

}
