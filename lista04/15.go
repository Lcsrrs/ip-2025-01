package main

import f "fmt"

func main() {
	var vetor1 []int
	var vetor2 []int
	var temp int

	for i := 0; i < 30; i++ {
		f.Printf("Digite o %dº elemento do vetor 1: ", i+1)
		f.Scan(&temp)
		vetor1 = append(vetor1, temp)
	}

	for i := 0; i < len(vetor1); i++ {
		if i%2 == 0 {
			vetor2 = append(vetor2, 2*vetor1[i])
		} else {
			vetor2 = append(vetor2, 3*vetor1[i])
		}
	}

	f.Print(vetor2)

}
