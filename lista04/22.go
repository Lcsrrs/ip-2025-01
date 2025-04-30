package main

import f "fmt"

func main() {
	var conta []int
	var saldoConta []float64
	var tempConta, codigo int
	var tempSaldo float64

	i := 0
	for i < 10 {
		f.Printf("Digite o código da %dª conta e o seu respectivo saldo: ", i+1)
		f.Scan(&tempConta, &tempSaldo)
		if i == 0 {
			conta = append(conta, tempConta)
			saldoConta = append(saldoConta, tempSaldo)
			i++
			continue
		}

		duplicada := false
		for j := 0; j < len(conta); j++ {
			if tempConta == conta[j] {
				f.Print("Conta já existente, tente novamente\n")
				duplicada = true
				break
			}
		}
		if !duplicada {
			conta = append(conta, tempConta)
			saldoConta = append(saldoConta, tempSaldo)
			i++
		}
	}

	f.Print("\n\n")

	for codigo != 4 {
		f.Print("1. Efetuar depósito\n2. Efetuar Saque\n3. Consultar ativo bancário\n4. Finalizar programa\n")
		f.Scan(&codigo)

		switch codigo {
		case 1:
			f.Print("Digite o código da conta e o valor a ser depositado: ")
			f.Scan(&tempConta, &tempSaldo)
			encontrou := false
			for i := 0; i < len(conta); i++ {
				if tempConta == conta[i] {
					encontrou = true
					saldoConta[i] += tempSaldo
					f.Printf("Conta %d atualizada, saldo atual: R$ %.2f\n\n", tempConta, saldoConta[i])
					codigo = 0
					break
				}
			}
			if !encontrou {
				f.Print("Conta não encontrada\n\n")
				codigo = 0
				break
			}

		case 2:
			f.Print("Digite o código da conta e o valor a ser sacado: ")
			f.Scan(&tempConta, &tempSaldo)
			encontrou := false
			for i := 0; i < len(conta); i++ {
				if tempConta == conta[i] {
					encontrou = true
					if saldoConta[i] < tempSaldo {
						f.Print("Saldo insuficiente\n\n")
						codigo = 0
						break
					} else {
						saldoConta[i] -= tempSaldo
						f.Printf("Conta %d atualizada, saldo atual: R$ %.2f\n\n", tempConta, saldoConta[i])
						codigo = 0
						break
					}
				}
			}
			if !encontrou {
				f.Print("Conta não encontrada\n\n")
				codigo = 0
				break
			}
		case 3:
			soma := 0.0
			for i := 0; i < len(saldoConta); i++ {
				soma += saldoConta[i]
			}
			f.Printf("Ativo bancário: R$ %.2f\n\n", soma)
		default:
			if codigo == 4 {
				f.Print("Programa finalizado\n\n")
			} else {
				f.Print("Opção inválida\n\n")
			}
		}
	}
}
