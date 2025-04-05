package main

import "fmt"

func main() {
	var sal_min, gasto float32
	fmt.Print("Valor do salário mínimo: ")
	fmt.Scan(&sal_min)
	fmt.Print("\nQuantidade de energia gasta na residência (kW): ")
	fmt.Scan(&gasto)

	var custokw = sal_min * 0.7 / 100
	var custoconsumo = custokw * gasto
	var custocomdesconto = custoconsumo * 0.9

	fmt.Printf("\nCusto por kW: %.2f", custokw)
	fmt.Printf("\nCusto do Consumo: %.2f", custoconsumo)
	fmt.Printf("\nCusto com desconto: %.2f", custocomdesconto)

}
