package main

import f "fmt"

func main() {
	var vetor []int
	var temp int
	var ant1 = 1
	var ant2 = 1

	vetor = append(vetor, ant1)
	vetor = append(vetor, ant2)

	for len(vetor) < 50 {
		vetor = append(vetor, ant1+ant2)
		temp = ant1
		ant1 = ant2
		ant2 += temp
	}

	f.Print(vetor)
}
