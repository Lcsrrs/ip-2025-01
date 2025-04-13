package main

import (
	f "fmt"
)

func main() {
	var n1, n2 int
	var quociente, resto int

	f.Print("Digite n1 e n2: ")
	f.Scan(&n1, &n2)
	resto = n1

	for resto >= n2 {
		resto -= n2
		quociente++
	}

	f.Printf("Quociente(%d,%d) = %d e o Resto(%d,%d)=%d", n1, n2, quociente, n1, n2, resto)

}
