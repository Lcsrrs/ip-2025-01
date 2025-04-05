package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)

	if idade <= 0 {
		fmt.Print("Idade inválida")
	} else if idade <= 2 {
		fmt.Print("Classificação: Recém Nascido")
	} else if idade <= 11 {
		fmt.Print("Classificação: Criança")
	} else if idade <= 19 {
		fmt.Print("Classificação: Adolescente")
	} else if idade <= 55 {
		fmt.Print("Classificação: Adulto")
	} else {
		fmt.Print("Classificação: Idoso")

	}
}
