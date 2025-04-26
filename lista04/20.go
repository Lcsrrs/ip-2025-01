package main

import f "fmt"

func main() {
	var jogadas []int
	type Repete struct {
		numeroRepetido   int
		numeroRepeticoes int
	}
	var repeticoes []Repete
	var temp Repete
	var temporario, i int

	for i < 20 {
		f.Printf("Digite a %dº jogada de dado: ", i+1)
		f.Scan(&temporario)
		if temporario > 6 || temporario < 1 {
			f.Print("Número inválido\n")
			continue
		} else {
			jogadas = append(jogadas, temporario)
			i++
		}
	}

	for i := 0; i < len(jogadas); i++ {
		for k := 0; k < len(repeticoes); k++ {
			if jogadas[i] == repeticoes[k].numeroRepetido {
				repeticoes[k].numeroRepeticoes++
				break
			} else if jogadas[i] != repeticoes[k].numeroRepetido && k == len(repeticoes)-1 {
				temp.numeroRepetido = jogadas[i]
				temp.numeroRepeticoes = 0
				repeticoes = append(repeticoes, temp)
			}
		}
		if len(repeticoes) == 0 {
			temp.numeroRepetido = jogadas[i]
			temp.numeroRepeticoes = 1
			repeticoes = append(repeticoes, temp)
		}
	}

	for i := 0; i < len(repeticoes); i++ {
		f.Printf("O valor %d repete %d vezes\n", repeticoes[i].numeroRepetido, repeticoes[i].numeroRepeticoes)
	}
}
