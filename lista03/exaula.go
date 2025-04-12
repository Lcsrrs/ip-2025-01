package main

import f "fmt"

func main() {

	var lista = []int{1, 2, 5, 7, 1, 9, 10, 1}
	var num int

	f.Print("Número a procurar: ")
	f.Scan(&num)

	for i := 0; i < len(lista); i++ {
		if lista[i] == num {
			f.Print("Índice: ", i, "\n")
		}
	}
}
