package main

import f "fmt"

func main() {
	var quantidade, temp, consecutivo, maior_consecutivo int
	var vetor []int

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		f.Scan(&temp)
		vetor = append(vetor, temp)
	}

	temp = vetor[0]
	consecutivo = 1

	for i := 1; i < len(vetor); i++ {
		if vetor[i] == temp+1 {
			consecutivo++
		} else {
			if consecutivo > maior_consecutivo {
				maior_consecutivo = consecutivo
			}
			consecutivo = 1
		}
		temp = vetor[i]
	}

	if maior_consecutivo == 0 {
		maior_consecutivo = consecutivo
	}

	f.Print(maior_consecutivo)
}
