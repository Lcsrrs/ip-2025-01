package main

import f "fmt"

func main() {
	var vetor []int
	var temp, menor, indMenor int

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor: ", i+1)
		f.Scan(&temp)
		vetor = append(vetor, temp)
	}

	menor = vetor[1]

	for i := 0; i < len(vetor); i++ {
		if vetor[i] < menor {
			menor = vetor[i]
			indMenor = i
		}
	}

	f.Printf("O menor elemento do vetor é %d e está na posição %d", menor, indMenor+1)
}
