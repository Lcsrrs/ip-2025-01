package main

import f "fmt"

func main() {
	var vetor []int
	var temp, quantidade int

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		f.Scan(&temp)
		vetor = append(vetor, temp)
		menor := vetor[0]
		for j := 0; j < len(vetor); j++ {
			if vetor[i] < vetor[j] {
				menor = vetor[i]
				vetor[i] = vetor[j]
				vetor[j] = menor
			}
		}
	}

	for i := 0; i < len(vetor); i++ {
		f.Print(vetor[i], " ")
	}
}
