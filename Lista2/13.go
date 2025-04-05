package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite um valor de 3 dígitos: ")
	fmt.Scan(&num)

	if num < 0 || float32(num)/100 < 1 || float32(num)/100 >= 10 {
		fmt.Print("Número inválido")
	} else {
		num /= 10
		fmt.Print("O dígito das dezenas é igual a ", num%10)
	}

}
