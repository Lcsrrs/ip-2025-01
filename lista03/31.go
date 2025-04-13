package main

import (
	f "fmt"
)

func main() {
	var total, n2 int

	for i := 0; i < 64; i++ {
		n2 = 1
		for j := 0; j < i; j++ {
			n2 *= 2
		}
		total += n2
	}
	f.Print("O total de grãos de milho será: ", total)
}
