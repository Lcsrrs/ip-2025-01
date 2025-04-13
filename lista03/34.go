package main

import (
	f "fmt"
)

func main() {
	var n1, n2, maior int

	f.Print("Digite n1 e n2: ")
	f.Scan(&n1, &n2)

	if n1 > n2 {
		maior = n1
	} else {
		maior = n2
	}

	for i := 2; i <= maior/2; i++ {
		if n1%i == 0 && n2%i == 0 {
			f.Printf("MMC %d e %d: %d", n1, n2, i)
			return
		}
	}
	f.Print("Não há MMC")
}
