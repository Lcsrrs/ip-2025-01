package main

import f "fmt"

func main() {
	var (
		vetor1     []int
		vetor2     []int
		vetorPar   []int
		vetorImpar []int
	)
	var temp, soma2 int

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do primeiro vetor: ", i+1)
		f.Scan(&temp)
		vetor1 = append(vetor1, temp)
	}

	for i := 0; i < 5; i++ {
		f.Printf("Digite o %dº elemento do segundo vetor: ", i+1)
		f.Scan(&temp)
		soma2 += temp
		vetor2 = append(vetor2, temp)
	}

	for i := 0; i < len(vetor1); i++ {
		if vetor1[i]%2 == 0 {
			vetorPar = append(vetorPar, vetor1[i]+soma2)
		} else {
			vetorImpar = append(vetorImpar, vetor1[i]+soma2)
		}
	}

	f.Print("Primeiro vetor resultante: ", vetorPar, "\n")
	f.Print("Segundo vetor resultante: ", vetorImpar)
}
