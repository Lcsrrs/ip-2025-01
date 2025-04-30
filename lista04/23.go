package main

import f "fmt"

func main() {
	janela := [24]int{}
	corredor := [24]int{}

	for i := 0; i < len(janela); i++ {
		janela[i] = 0
		corredor[i] = 0
	}

	var escolha int

	for true {
		var somaJanela, somaCorredor int

		for i := 0; i < len(corredor); i++ {
			if janela[i] == 1 {
				somaJanela++
			}
			if corredor[i] == 1 {
				somaCorredor++
			}
		}

		if somaCorredor == len(corredor) && somaJanela == len(janela) {
			f.Print("Todos os assentos estão ocupados\n3 - Sair\n")
			f.Scan(&escolha)
			if escolha != 3 {
				f.Print("Escolha inválida\n\n")
				continue
			}
		} else if somaCorredor == len(corredor) {
			f.Print("Os assentos do corredor estão esgotados.\n2 - Janela\n3 - Sair\n")
			f.Scan(&escolha)
			if escolha != 2 && escolha != 3 {
				f.Print("Escolha inválida\n\n")
				continue
			}
		} else if somaJanela == len(janela) {
			f.Print("Os assentos da janela estão esgotados.\n1 - Corredor\n3 - Sair\n")
			f.Scan(&escolha)
			if escolha != 1 && escolha != 3 {
				f.Print("Escolha inválida\n\n")
				continue
			}
		} else {
			f.Print("Deseja escolher poltrona no corredor ou na janela?\n1 - Corredor\n2 - Janela\n3- Sair\n")
			f.Scan(&escolha)
			if escolha != 1 && escolha != 2 && escolha != 3 {
				f.Print("Escolha inválida\n\n")
				continue
			}
		}
		if escolha == 3 {
			return
		}

		if escolha == 1 {
			var livres []int
			var assento int
			for i := 0; i < len(corredor); i++ {
				if corredor[i] == 0 {
					livres = append(livres, i+1)
				}
			}
			disponivel := false
			for !disponivel {
				if len(livres) == len(corredor) {
					f.Print("Todas as poltronas do corredor estão vazias, qual deseja escolher? (1-24)\n")
				} else {
					f.Print("Somente as poltronas ", livres, " do corredor estão vazias, qual deseja escolher?\n")
				}
				f.Scan(&assento)
				for i := 0; i < len(livres); i++ {
					if assento == livres[i] {
						disponivel = true
						corredor[livres[i]-1] = 1
						f.Print("Assento reservado\n\n")
					} else if i == len(livres)-1 && !disponivel {
						f.Print("Assento inválido, digite outro valor\n")
					}
				}
			}

		}

		if escolha == 2 {
			var livres []int
			var assento int
			for i := 0; i < len(janela); i++ {
				if janela[i] == 0 {
					livres = append(livres, i+1)
				}
			}
			disponivel := false
			for !disponivel {
				if len(livres) == len(janela) {
					f.Print("Todas as poltronas da janela estão vazias, qual deseja escolher? (1-24)\n")
				} else {
					f.Print("Somente as poltronas ", livres, " da janela estão vazias, qual deseja escolher?\n")
				}
				f.Scan(&assento)
				for i := 0; i < len(livres); i++ {
					if assento == livres[i] {
						disponivel = true
						janela[livres[i]-1] = 1
						f.Print("Assento reservado\n\n")
					} else if i == len(livres)-1 && !disponivel {
						f.Print("Assento inválido, digite outro valor\n")
					}
				}
			}

		}
	}

}
