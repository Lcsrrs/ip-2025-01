package main

import (
	f "fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var bois []int
	var maisGordo, indGordo, indMagro int
	maisMagro := 1000

	for i := 0; i < 90; i++ {
		bois = append(bois, randInt(100, 1000))
	}

	for i := 0; i < len(bois); i++ {
		if bois[i] > maisGordo {
			maisGordo = bois[i]
			indGordo = i
		}
		if bois[i] < maisMagro {
			maisMagro = bois[i]
			indMagro = i
		}
	}

	f.Printf("O boi mais pesado é o de índice %d, com o peso %d kg, e o mais mais magro é o de índice %d, com o peso de %d kg", indGordo, maisGordo, indMagro, maisMagro)
}
