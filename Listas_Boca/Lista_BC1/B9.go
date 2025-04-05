package main

import "fmt"

func main() {
	var peso, altura, imc float64

	fmt.Scan(&peso, &altura)

	imc = peso / (altura * altura)

	if imc < 18.5 {
		fmt.Print("Abaixo do peso")
	} else if imc < 25 {
		fmt.Print("Peso normal")
	} else if imc < 30 {
		fmt.Print("Sobrepeso")
	} else {
		fmt.Print("Obesidade")
	}
}
