package main

import f "fmt"

func main() {
	var N int

	f.Print("Forneça o valor de N: ")
	f.Scan(&N)

	for i := 1; i <= N; i++ {
		f.Print(i*i, " ")
	}
}
