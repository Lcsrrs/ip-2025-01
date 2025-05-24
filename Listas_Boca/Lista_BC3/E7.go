package main

import f "fmt"

func main() {
	var quantidade, temp, produto int
	var vetor1 []int
	var vetor2 []int

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		f.Scan(&temp)
		vetor1 = append(vetor1, temp)
	}

	for i := 0; i < quantidade; i++ {
		f.Scan(&temp)
		vetor2 = append(vetor2, temp)
		produto += vetor1[i] * vetor2[i]
	}

	f.Print(produto)
}
