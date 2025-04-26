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

	for i := 0; i < 15; i++ {
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

	for i := 0; i <= 10; i++ {
		achou := false
		for j := 0; j < len(repeticoes); j++ {
			if repeticoes[j].numeroRepetido == i {
				achou = true
				f.Printf("Nota %d -> FA=%d    FR=%.2f %%\n", i, repeticoes[j].numeroRepeticoes, float64(repeticoes[j].numeroRepeticoes*100)/float64(len(vetor)))
			} else if j == len(repeticoes)-1 && achou == false {
				f.Printf("Nota %d -> FA=0    FR=0 %%\n", i)
			}
		}
	}
}
