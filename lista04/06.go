package main

import f "fmt"

func main() {
	var vetor []int

	for i := 100; i > 0; i-- {
		vetor = append(vetor, i)
	}

	f.Print(vetor)
}
