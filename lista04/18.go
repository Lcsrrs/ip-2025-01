package main

import f "fmt"

func main() {
	var vetor []int
	var temp int

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor: ", i+1)
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

	f.Print(vetor)
}
