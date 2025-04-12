package main

import f "fmt"

func main() {
	var d int
	var soma float64

	f.Print("Digite o número de termos: ")
	f.Scan(&d)

	if d <= 0 {
		f.Print("Número de termos inválido")
		return
	}

	for i := 0; i < d; i++ {
		if d%2 == 0 {
			soma -= (1000.0 - 3.0*float64(i)) / float64(d)
		} else {
			soma += (1000.0 - 3.0*float64(i)) / float64(d)
		}
	}

	f.Print("Soma final = ", soma)
}
