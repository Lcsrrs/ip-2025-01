package main

import f "fmt"

func main() {
	var vetor []int
	type Repete struct {
		numeroRepetido   int
		numeroRepeticoes int
	}
	var repeticoes []Repete
	var temp Repete
	var temporario, quantidade int

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		f.Scan(&temporario)
		vetor = append(vetor, temporario)
	}

	for i := 0; i < len(vetor); i++ {
		for k := 0; k < len(repeticoes); k++ {
			if vetor[i] == repeticoes[k].numeroRepetido {
				repeticoes[k].numeroRepeticoes++
				break
			} else if vetor[i] != repeticoes[k].numeroRepetido && k == len(repeticoes)-1 {
				temp.numeroRepetido = vetor[i]
				temp.numeroRepeticoes = 0
				repeticoes = append(repeticoes, temp)
			}
		}
		if len(repeticoes) == 0 {
			temp.numeroRepetido = vetor[i]
			temp.numeroRepeticoes = 1
			repeticoes = append(repeticoes, temp)
		}
	}

	f.Print(len(repeticoes))
}
