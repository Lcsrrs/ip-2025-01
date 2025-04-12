package main

import f "fmt"

func main() {
	var N int
	fatorial := 1

	f.Print("Digite o valor: ")
	f.Scan(&N)

	if N <= 0 {
		f.Print("Número inválido")
		return
	}

	for i := 0; i < N; i++ {
		fatorial *= (N - i)
	}

	f.Printf("Fatorial do número %d igual a %d", N, fatorial)
}
