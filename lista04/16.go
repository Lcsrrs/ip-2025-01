package main

import (
	f "fmt"
	"math/rand/v2"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func main() {
	var vetor []int
	type Repete struct {
		idadeRepetida    int
		numeroRepeticoes int
	}
	var repeticoes []Repete
	var temp Repete
	var moda, maior int

	for i := 0; i < 50; i++ {
		vetor = append(vetor, randRange(0, 10))
	}

	for i := 0; i < len(vetor); i++ {
		for k := 0; k < len(repeticoes); k++ {
			if vetor[i] == repeticoes[k].idadeRepetida {
				repeticoes[k].numeroRepeticoes++
				break
			} else if vetor[i] != repeticoes[k].idadeRepetida && k == len(repeticoes)-1 {
				temp.idadeRepetida = vetor[i]
				temp.numeroRepeticoes = 0
				repeticoes = append(repeticoes, temp)
			}
		}
		if len(repeticoes) == 0 {
			temp.idadeRepetida = vetor[i]
			temp.numeroRepeticoes = 1
			repeticoes = append(repeticoes, temp)
		}
	}

	for i := 0; i < len(repeticoes); i++ {
		if repeticoes[i].numeroRepeticoes > maior {
			maior = repeticoes[i].numeroRepeticoes
			moda = repeticoes[i].idadeRepetida
		}
	}

	f.Print(vetor, "\n")
	f.Print("Moda: ", moda)
}
