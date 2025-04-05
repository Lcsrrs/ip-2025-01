package main

import "fmt"

func main() {
	var x1, x2, x3, menor, inter, maior int

	fmt.Scan(&x1, &x2, &x3)

	if x1 >= x2 && x1 >= x3 {
		maior = x1
		if x2 >= x3 {
			inter = x2
			menor = x3
		} else {
			inter = x3
			menor = x2
		}
	} else if x2 >= x1 && x2 >= x3 {
		maior = x2
		if x1 >= x3 {
			inter = x1
			menor = x3
		} else {
			inter = x3
			menor = x1
		}
	} else if x3 >= x1 && x3 >= x2 {
		maior = x3
		if x1 >= x2 {
			inter = x1
			menor = x2
		} else {
			inter = x2
			menor = x1
		}
	}

	fmt.Print(menor, inter, maior)
}
