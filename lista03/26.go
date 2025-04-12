package main

import f "fmt"

func main() {
	var soma float64
	var iFat int

	for i := 0; i <= 20; i++ {
		//Determinando i fatorial:
		iFat = 1
		for j := 0; j < i; j++ {
			iFat *= (i - j)
		}
		soma += (100 - float64(i)) / float64(iFat)
	}

	f.Print("Soma final: ", soma)
}
