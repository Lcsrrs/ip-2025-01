package main

import f "fmt"

func main() {
	var codigo int

	for true {
		f.Print("Digite o código:\n0 - Finalizar Programa\n1 - Vetor na ordem direta\n2 - Vetor na ordem inversa\n")
		f.Scan(&codigo)
		if codigo == 0 {
			return
		} else if codigo == 1 || codigo == 2 {
			break
		} else {
			f.Print("Código inválido, tente novamente\n")
		}
	}

	var vetor []float64
	var temp float64

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor ", i+1)
		f.Scan(&temp)
		vetor = append(vetor, temp)
	}

	f.Print("\n")
	if codigo == 1 {
		for i := 0; i < len(vetor); i++ {
			f.Print(vetor[i], "\n")
		}
	} else {
		for i := len(vetor) - 1; i >= 0; i-- {
			f.Print(vetor[i], "\n")
		}
	}

}
