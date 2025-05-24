package main

import f "fmt"

func main() {
	var linhas, colunas int
	var matriz [][]int
	f.Scan(&linhas, &colunas)
	matriz = make([][]int, linhas)
	for i := range matriz {
		matriz[i] = make([]int, colunas)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			f.Scan(&matriz[i][j])
		}
	}

	for i := 0; i < colunas; i++ {
		for j := 0; j < linhas; j++ {
			f.Print(matriz[j][i], " ")
		}
		f.Print("\n")
	}
}
