package main

import f "fmt"

func main() {
	d := 1
	var soma float64

	for i := 38; i > 1; i-- {
		soma += float64(i) * float64(i-1) / float64(d)
		d++
	}

	f.Print("Valor final da soma: ", soma)
}
