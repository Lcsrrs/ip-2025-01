package main

import f "fmt"

func main() {
	var vetor []int
	i := 1

	for len(vetor) < 100 {
		vetor = append(vetor, i)
		i += 2
	}

	f.Print(vetor)
}
