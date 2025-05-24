package main

import f "fmt"

func main() {
	var linhas, colunas int
	var matriz [][]int
	f.Scan(&linhas)
	colunas = linhas
	matriz = make([][]int, linhas)
	for i := range matriz {
		matriz[i] = make([]int, colunas)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			f.Scan(&matriz[i][j])
		}
	}

	for i := 0; i < len(matriz); i++ {
		for j := len(matriz) - 1; j >= 0; j-- {
			f.Print(matriz[j][i], " ")
		}
		f.Print("\n")
	}
}
