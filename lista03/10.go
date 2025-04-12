package main

import f "fmt"

func main() {
	var N, temp int
	anterior := 1
	anterior2 := 0

	f.Print("Quanitdade de itens: ")
	f.Scan(&N)

	f.Print("0-1")

	for i := 0; i < N; i++ {
		f.Print("-", anterior+anterior2)
		temp = anterior2
		anterior2 = anterior
		anterior = anterior + temp
	}
}
