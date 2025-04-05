package main

import "fmt"

func main() {
	var valor float32
	var opcao string

	fmt.Print("Digite o valor inicial do veículo: ")
	fmt.Scan(&valor)

	fmt.Print("O carro terá ar condicionado? +R$ 1.750,00 (S ou N): ")
	fmt.Scan(&opcao)
	if opcao == "S" || opcao == "s" {
		valor += 1750
	}
	opcao = "n"
	fmt.Print("O carro terá pintura metálica? +R$ 800,00 (S ou N): ")
	fmt.Scan(&opcao)
	if opcao == "S" || opcao == "s" {
		valor += 800
	}
	opcao = "n"
	fmt.Print("O carro terá vidro elétrico? +R$ 1.200,00 (S ou N): ")
	fmt.Scan(&opcao)
	if opcao == "S" || opcao == "s" {
		valor += 1200
	}
	opcao = "n"
	fmt.Print("O carro terá direção hidráulica? +R$ 2.000,00 (S ou N): ")
	fmt.Scan(&opcao)
	if opcao == "S" || opcao == "s" {
		valor += 2000
	}

	fmt.Print("O valor final do veículo é: R$", valor)
}
