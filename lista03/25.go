package main

import f "fmt"

func main() {
	var soma float64
	d := 15

	for i := 1; i < 15; i++ {
		if i%2 == 0 {
			soma -= float64(i) * float64(i) / float64(d) * float64(d)
			d--
		} else {
			soma -= float64(i) * float64(i) / float64(d) * float64(d)
			d--
		}
	}

	f.Print("Valor final da soma: ", soma)
}
