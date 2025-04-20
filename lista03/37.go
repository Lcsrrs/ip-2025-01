package main

import (
	f "fmt"
	m "math"
)

func main() {
	var base8, base8init int
	var base10, cont int

	f.Print("Digite o valor na base 8: ")
	f.Scan(&base8)
	base8init = base8

	for base8 > 0 {
		if base8%10 == 9 || base8%10 == 8 {
			f.Print("O número informado não está na base 8")
			return
		}
		base10 += base8 % 10 * int(m.Pow(8, float64(cont)))
		cont++
		base8 /= 10
	}

	f.Printf("O número %d na base 10 é igual a %d", base8init, base10)
}
