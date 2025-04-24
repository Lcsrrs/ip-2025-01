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
	var temporario int

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor: ", i+1)
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

	for i := 0; i < len(repeticoes); i++ {
		if repeticoes[i].numeroRepeticoes > 1 {
			f.Printf("O valor %d repete %d vezes\n", repeticoes[i].numeroRepetido, repeticoes[i].numeroRepeticoes)
		}
	}
}
