<<<<<<< HEAD
package main

import "fmt"

func main() {
	var valor float64
	var dia int
	var categoria string

	fmt.Print("Digite o valor de locação , o dia e a categoria do DVD - Comum(C), Lançamento(L): ")
	fmt.Scan(&valor, &dia, &categoria)

	if categoria == "C" || categoria == "L" {
		if categoria == "L" {
			valor *= 1.15
		}
	} else {
		fmt.Print("Categoria inválida")
		return
	}

	if dia == 2 || dia == 3 || dia == 5 {
		fmt.Print("Valor de locação: ", valor*0.6)
	} else {
		fmt.Print("Valor de locação: ", valor)
	}
}
=======
package main

import "fmt"

func main() {
	var valor float64
	var dia int
	var categoria string

	fmt.Print("Digite o valor de locação , o dia e a categoria do DVD - Comum(C), Lançamento(L): ")
	fmt.Scan(&valor, &dia, &categoria)

	if categoria == "C" || categoria == "L" {
		if categoria == "L" {
			valor *= 1.15
		}
	} else {
		fmt.Print("Categoria inválida")
		return
	}

	if dia == 2 || dia == 3 || dia == 5 {
		fmt.Print("Valor de locação: ", valor*0.6)
	} else {
		fmt.Print("Valor de locação: ", valor)
	}
}
>>>>>>> bf15460b564febf0a3bad8ba55fdcf48a890b0ef
