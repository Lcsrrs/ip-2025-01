package main

import "fmt"

func main() {
	var (
		conta, consumo, valor_conta float32
	)
	var tipo string

	fmt.Print("Escreva o número da conta, consumo em m³ e o tipo de cobrança (R, C ou I): ")
	fmt.Scan(&conta, &consumo, &tipo)
	fmt.Print("CONTA: ", conta)
	fmt.Print("\n")
	if tipo == "R" {
		valor_conta = 5 + consumo*0.05
		fmt.Print("VALOR CONTA: ", valor_conta)
	} else if tipo == "C" {
		if consumo <= 80 {
			fmt.Print("VALOR CONTA: 500")
		} else {
			valor_conta = 500 + (consumo-80)*0.25
			fmt.Print("VALOR CONTA: ", valor_conta)
		}
	} else if tipo == "I" {
		if consumo <= 100 {
			fmt.Print("VALOR CONTA: 800")
		} else {
			valor_conta = 800 + (consumo-100)*0.04
			fmt.Print("VALOR CONTA: ", valor_conta)
		}
	} else {
		fmt.Print("Tipo não conhecido")
	}

}
