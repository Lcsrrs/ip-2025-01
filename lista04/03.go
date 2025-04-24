package main

import f "fmt"

func main() {
	var vetor []int
	var temp, quantPar, quantImpar int

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor: ", i+1)
		f.Scan(&temp)
		vetor = append(vetor, temp)
	}

	f.Print("a) Números pares do vetor: ")

	for i := 0; i < len(vetor); i++ {
		if vetor[i]%2 == 0 {
			f.Print(" ", vetor[i], " ")
			quantPar++
		}
	}
	f.Print("\nb) Quantidade de números pares: ", quantPar)
	f.Print("\nc) Números ímpares do vetor: ")
	for i := 0; i < len(vetor); i++ {
		if vetor[i]%2 != 0 {
			f.Print(" ", vetor[i], " ")
			quantImpar++
		}
	}
	f.Print("\nd) Quantidade de números ímpares: ", quantImpar)
}
