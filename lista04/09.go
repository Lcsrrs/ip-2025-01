package main

import f "fmt"

func main() {
	var vetor []float64
	var temp, somaAltura, media float64

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor: ", i+1)
		f.Scan(&temp)
		somaAltura += temp
		vetor = append(vetor, temp)
	}

	media = somaAltura / float64(len(vetor))

	for i := 0; i < len(vetor); i++ {
		if vetor[i] > media {
			f.Print(vetor[i], "\n")
		}
	}

}
