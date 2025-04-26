package main

import f "fmt"

func main() {
	var vetor1 []int
	var vetor2 []int
	var temp int

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor 1: ", i+1)
		f.Scan(&temp)
		vetor1 = append(vetor1, temp)
	}
	for i := 0; i < 5; i++ {
		f.Printf("Digite o %dº elemento do vetor 2: ", i+1)
		f.Scan(&temp)
		vetor2 = append(vetor2, temp)
	}

	for i := 0; i < len(vetor1); i++ {
		f.Print("Numero ", vetor1[i], "\n")
		divisivel := false
		for j := 0; j < len(vetor2); j++ {
			if vetor1[i]%vetor2[j] == 0 {
				f.Printf("	Divisível por %d na posição %d\n", vetor2[j], j)
				divisivel = true
			} else if j == len(vetor2)-1 && divisivel == false {
				f.Print("	Não é divisível por nenhum valor\n")
			}
		}
	}
}
