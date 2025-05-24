package main

import f "fmt"

func main() {
	var vetor []float64
	var temp, somaAltura, media float64
	var quantidade int

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
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
