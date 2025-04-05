package main

import "fmt"

func main() {

	var tempF, tempC, chuvaPol, chuvaMm float32

	fmt.Print("Escreva a temperatura em Fahrenheit e a precipitação em polegadas: ")
	fmt.Scan(&tempF, &chuvaPol)

	tempC = 5 * (tempF - 32) / 9
	chuvaMm = 25.4 * chuvaPol

	fmt.Print("O VALOR EM CELSIUS: ", tempC, "\nA QUANTIDADE DE CHUVA É: ", chuvaMm)

}
