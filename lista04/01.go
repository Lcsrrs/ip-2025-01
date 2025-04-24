package main

import f "fmt"

func main() {
	var vetor []int
	var temp int

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor: ", i+1)
		f.Scan(&temp)
		vetor = append(vetor, temp)
	}

	for i := 0; i < len(vetor); i++ {
		if vetor[i] > 50 {
			f.Printf("O elemento %d do vetor é maior que 50 e tem o valor %d\n", i, vetor[i])
		}
	}

}
