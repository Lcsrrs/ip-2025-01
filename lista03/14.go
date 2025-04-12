package main

import f "fmt"

func main() {
	var N1, N2, contPrimos int

	f.Print("Forneça os valores N1 e N2: ")
	f.Scan(&N1, &N2)

	for i := N1; i < N2; i++ {
		if i == 2 {
			contPrimos++
			continue
		}
		for j := 2; j < i; j++ {
			if i%j == 0 {

				break
			} else if j == i-1 {

				contPrimos++
			}
		}
	}

	f.Printf("Entre os valores %d e %d existem %d números primos", N1, N2, contPrimos)
}
