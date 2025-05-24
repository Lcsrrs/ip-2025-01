package main

import f "fmt"

func main() {
	var quantidade int
	var temp, maior, menor float64
	var vetor []float64

	f.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		f.Scan(&temp)
		vetor = append(vetor, temp)
		if i == 0 {
			maior = temp
			menor = temp
			continue
		}

		if temp > maior {
			maior = temp
		}

		if temp < menor {
			menor = temp
		}
	}

	if maior == menor {
		for i := 0; i < quantidade; i++ {
			f.Print("0.00\n")
		}
	} else {
		for i := 0; i < quantidade; i++ {
			f.Printf("%.2f\n", (vetor[i]-menor)/(maior-menor))
		}
	}

}
