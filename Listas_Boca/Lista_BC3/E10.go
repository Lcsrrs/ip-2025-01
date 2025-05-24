package main

import f "fmt"

func main() {
	var quantidade, temp int
	var vetor []int

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		f.Scan(&temp)
		vetor = append(vetor, temp)
	}

	palindromo := true

	if quantidade%2 == 0 {
		for i := 0; i < quantidade/2; i++ {
			if vetor[i] != vetor[quantidade-i-1] {
				palindromo = false
				break
			}
		}
	} else {
		for i := 0; i < (quantidade-1)/2; i++ {
			if vetor[i] != vetor[quantidade-i-1] {
				palindromo = false
				break
			}
		}
	}

	if palindromo {
		f.Print("SIM")
	} else {
		f.Print("NAO")
	}

}
