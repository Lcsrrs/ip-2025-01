package main

import "fmt"

func main() {

	var horas, valor int

	fmt.Print("Escreva a quantidade de horas: ")
	fmt.Scan(&horas)

	if horas%3 == 0 {
		valor = 10 * (horas / 3)
	} else if horas%3 == 1 {
		valor = 10*(horas/3) + 5
	} else {
		valor = 10*(horas/3) + 10
	}

	fmt.Print("Valor: ", valor)

	//fmt.Print("O VALOR A PAGAR É: ", valor, "\nPrimeiro termo: ", 10*(horas/3+(3-horas%3)), "\nSegundo termo: ", 5*(3-horas%3))

}
