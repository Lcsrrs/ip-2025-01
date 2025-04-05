package main

import "fmt"

func main() {
	var quantidade int

	fmt.Print("Escreva a quantidade: ")
	fmt.Scan(&quantidade)

	if quantidade < 5 {
		fmt.Print("Quantidade muito baixa")
		return
	} else if quantidade > 2000 {
		fmt.Print("Quantidade muito alta")
		return
	} else {
		for i := 2; i <= quantidade; i += 2 {
			fmt.Print(i, "^2: ", i*i, "\n")
		}
	}
}
