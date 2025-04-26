package main

import f "fmt"

func main() {
	var vetor []int
	var primos []int
	var temp int

	for i := 0; i < 10; i++ {
		f.Printf("Digite o %dº elemento do vetor: ", i+1)
		f.Scan(&temp)
		vetor = append(vetor, temp)
	}

	for i := 0; i < len(vetor); i++ {
		if vetor[i] == 1 || vetor[i] == 2 {
			primos = append(primos, vetor[i])
			primos = append(primos, i)
			continue
		}
		primo := true
		for j := 2; j < vetor[i]; j++ {
			if vetor[i]%j == 0 {
				primo = false
				break
			} else if j == vetor[i]-1 && primo == true {
				primos = append(primos, vetor[i])
				primos = append(primos, i)
			}
		}
	}

	f.Print("Valores primos: \n")
	for i := 0; i < len(primos); i += 2 {
		f.Print(primos[i], " na posição ", primos[i+1]+1, "\n")
	}
}
