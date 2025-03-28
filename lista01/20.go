package main

import "fmt"

func main() {
	var hora, minuto, segundo, segundoTotal int

	fmt.Print("Escreva o tempo separado por espaço: ")
	fmt.Scan(&hora, &minuto, &segundo)

	segundoTotal = hora*3600 + minuto*60 + segundo

	fmt.Print("O tempo total em segundos é: ", segundoTotal)

}
