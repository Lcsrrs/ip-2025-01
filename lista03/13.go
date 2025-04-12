package main

import f "fmt"

func main() {
	var H float64
	var den = 1

	for i := 1; i < 100; i += 2 {
		H += float64(i) / float64(den)
		den++
	}

	f.Print("O valor de H é: ", H)
}
