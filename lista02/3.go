package main

import "fmt"

func main() {
	var x1, x2 int

	fmt.Print("Digite dois números inteiros: ")
	fmt.Scan(&x1, &x2)

	if x1+x2 > 20 {
		fmt.Print("O valor somado é maior que 20, e esse valor mais 8 é igual a ", x1+x2+8)
	} else {
		fmt.Print("O valor somado é menor que 20, e esse valor menos 5 é igual a ", x1+x2-5)
	}
}
