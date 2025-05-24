package main

import f "fmt"

func main() {
	var quantidade, temp int
	var vetor []int
	var vetor_resultante []int

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		f.Scan(&temp)
		vetor = append(vetor, temp)
	}

	if quantidade == 1 {
		vetor_resultante = append(vetor_resultante, 0)
	} else {
		for i := 0; i < len(vetor); i++ {
			if i == 0 {
				vetor_resultante = append(vetor_resultante, vetor[i+1])
			} else if i == len(vetor)-1 {
				vetor_resultante = append(vetor_resultante, vetor[i-1])
			} else {
				vetor_resultante = append(vetor_resultante, vetor[i-1]+vetor[i+1])
			}
		}
	}

	for i := 0; i < len(vetor_resultante); i++ {
		f.Print(vetor_resultante[i], " ")
	}
}
