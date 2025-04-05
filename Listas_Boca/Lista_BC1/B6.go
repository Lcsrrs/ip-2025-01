package main

import "fmt"

func main() {
	var conta int
	var limite, saldo, depositos, retiradas float64

	fmt.Scan(&conta, &limite, &saldo, &depositos, &retiradas)

	saldo_final := saldo + depositos - retiradas

	if saldo_final > limite {
		fmt.Print("Credito excedido: ", saldo_final)
	} else {
		fmt.Print("Saldo: ", saldo_final)
	}
}
