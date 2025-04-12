package main

import f "fmt"

func main() {
	var soma int

	for i := 1; i <= 20; i++ {
		soma += i
		f.Print(i, "\n")
	}

	f.Print("Soma final: ", soma)
}
