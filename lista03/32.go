package main

import (
	f "fmt"
)

func main() {
	var n1, n2, mult int

	f.Print("Digite n1 e n2: ")
	f.Scan(&n1, &n2)
	mult = n1

	for i := 1; i < n2; i++ {
		mult += n1
	}

	f.Printf("%d x %d = %d", n1, n2, mult)
}
