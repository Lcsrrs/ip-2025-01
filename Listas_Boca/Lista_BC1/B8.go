package main

import "fmt"

func main() {
	var vendas float64

	fmt.Scan(&vendas)

	if vendas <= 15000 {
		fmt.Print(500 + vendas*0.09)
	} else {
		fmt.Print(500 + 800 + vendas*0.09)
	}
}
