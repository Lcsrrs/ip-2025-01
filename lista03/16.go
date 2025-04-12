package main

import f "fmt"

func main() {
	var a1, a2, N, ai, temp int

	f.Print("Digite os dois primeiros termos da série e a quantidade de termos: ")
	f.Scan(&a1, &a2, &N)

	f.Printf("Série de Fetuccine:\n%d\n%d\n", a1, a2)

	for i := 3; i <= N; i++ {
		if i%2 != 0 {
			ai = a2 + a1
			f.Print(ai, "\n")
			temp = a2
			a2 = ai
			a1 = temp
		} else {
			ai = a2 - a1
			f.Print(ai, "\n")
			temp = a2
			a2 = ai
			a1 = temp
		}
	}
}
