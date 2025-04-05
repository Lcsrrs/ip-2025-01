<<<<<<< HEAD
package main

import "fmt"

func main() {
	var conta int
	var gasto, valor_final float64
	var tipo string

	fmt.Print("Digite o número da conta, o gasto em m³ e seu tipo(R, C ou I): ")
	fmt.Scan(&conta, &gasto, &tipo)

	fmt.Print("Conta: ", conta, "\n")
	if tipo == "R" {
		valor_final = 5 + 0.05*gasto
		fmt.Printf("Valor da conta: R$ %.2f \n", valor_final)
	} else if tipo == "C" {
		if gasto < 80 {
			fmt.Print("Valor da conta: R$ 500,00")
		} else {
			valor_final = 500 + (gasto-80)*0.25
			fmt.Printf("Valor da conta: R$ %.2f \n", valor_final)
		}
	} else if tipo == "I" {
		if gasto < 100 {
			fmt.Print("Valor da conta: R$ 800,00")
		} else {
			valor_final = 800 + (gasto-100)*0.04
			fmt.Printf("Valor da conta: R$ %.2f \n", valor_final)
		}

	} else {
		fmt.Print("Tipo inválido")
	}
}
=======
package main

import "fmt"

func main() {
	var conta int
	var gasto, valor_final float64
	var tipo string

	fmt.Print("Digite o número da conta, o gasto em m³ e seu tipo(R, C ou I): ")
	fmt.Scan(&conta, &gasto, &tipo)

	fmt.Print("Conta: ", conta, "\n")
	if tipo == "R" {
		valor_final = 5 + 0.05*gasto
		fmt.Printf("Valor da conta: R$ %.2f \n", valor_final)
	} else if tipo == "C" {
		if gasto < 80 {
			fmt.Print("Valor da conta: R$ 500,00")
		} else {
			valor_final = 500 + (gasto-80)*0.25
			fmt.Printf("Valor da conta: R$ %.2f \n", valor_final)
		}
	} else if tipo == "I" {
		if gasto < 100 {
			fmt.Print("Valor da conta: R$ 800,00")
		} else {
			valor_final = 800 + (gasto-100)*0.04
			fmt.Printf("Valor da conta: R$ %.2f \n", valor_final)
		}

	} else {
		fmt.Print("Tipo inválido")
	}
}
>>>>>>> bf15460b564febf0a3bad8ba55fdcf48a890b0ef
