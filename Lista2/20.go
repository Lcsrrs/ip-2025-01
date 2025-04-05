package main

import "fmt"

func main() {
	var preco float64
	var tipo int

	fmt.Print("Digite o valor e o tipo de pagamento, sendo:\n1 - À vista, dinheiro ou cheque, 10% de desconto\n2 - À vista, cartão de crédito, 5% de desconto\n3 - Em 2 vezes, preço normal de etiqueta sem juros\n4 - Em 3 vezes, preço normal de etiqueta + 10% de juros\n")
	fmt.Scan(&preco, &tipo)

	if tipo != 1 && tipo != 2 && tipo != 3 && tipo != 4 {
		fmt.Print("Tipo inválido")
		return
	}

	if tipo == 1 {
		fmt.Printf("Preço final: R$ %.2f", preco*0.9)
	} else if tipo == 2 {
		fmt.Printf("Preço final: R$ %.2f", preco*0.95)
	} else if tipo == 3 {
		fmt.Printf("Preço final: R$ %.2f", preco)
	} else {
		fmt.Printf("Preço final: R$ %.2f", preco*1.1)
	}

}
