package main

import f "fmt"

func main() {
	var quantidade, temp, soma int
	var vetor1 []int
	var vetor2 []int

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		f.Scan(&temp)
		vetor1 = append(vetor1, temp)
	}

	for i := 0; i < len(vetor1); i++ {
		if i == 0 {
			vetor2 = append(vetor2, vetor1[i])
		} else {
			for j := 0; j <= i; j++ {
				soma += vetor1[j]
			}
			vetor2 = append(vetor2, soma)
			soma = 0
		}
	}

	for i := 0; i < len(vetor2); i++ {
		f.Print(vetor2[i], " ")
	}
}
