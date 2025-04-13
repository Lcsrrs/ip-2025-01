package main

import f "fmt"

func main() {
	var n, soma int
	f.Scan(&n)

	for i := 1; i <= n; i++ {
		soma += i
		f.Print(i, " ")
	}
	f.Print("\n", soma)
}
