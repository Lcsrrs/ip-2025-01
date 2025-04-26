package main

import f "fmt"

func main() {
	var vetor1 []int
	var vetor2 []int
	var vetorResultante []int
	var temp int

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor 1: ", i+1)
		f.Scan(&temp)
		vetor1 = append(vetor1, temp)
	}

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor 2: ", i+1)
		f.Scan(&temp)
		vetor2 = append(vetor2, temp)
	}

	for i := 0; i < len(vetor1)+len(vetor2); i++ {
		if i%2 != 0 {
			vetorResultante = append(vetorResultante, vetor2[(i-1)/2])
		} else {
			vetorResultante = append(vetorResultante, vetor1[i/2])
		}
	}

	f.Print(vetorResultante)

}
