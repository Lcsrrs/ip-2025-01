package main

import "fmt"

func main() {
	var x1 float32

	fmt.Print("Digite o valor de compra: ")
	fmt.Scan(&x1)

	if x1 <= 0 {
		fmt.Print("Valor de compra inválido")
	} else if x1 < 10 {
		fmt.Print("Valor de venda é ", x1*1.7)
	} else if x1 < 30 {
		fmt.Print("Valor de venda é ", x1*1.5)
	} else if x1 < 50 {
		fmt.Print("Valor de venda é ", x1*1.4)
	} else {
		fmt.Print("Valor de venda é ", x1*1.3)
	}
}
