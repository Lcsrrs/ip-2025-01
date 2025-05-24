package main

import (
	f "fmt"
)

type error interface {
	Error() string
}

func main() {
	var linhas, colunas, tempSoma int
	var matriz [][]int
	var matriz_resultante [][]int
	f.Scan(&linhas, &colunas)
	matriz = make([][]int, linhas)
	matriz_resultante = make([][]int, linhas)
	for i := range matriz {
		matriz[i] = make([]int, colunas)
		matriz_resultante[i] = make([]int, colunas)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			f.Scan(&matriz[i][j])
		}
	}

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			tempSoma = 0
			for k := -1; k < 2; k++ {
				for l := -1; l < 2; l++ {
					if k != 0 || l != 0 {
						if (i+k) >= 0 && (i+k) < len(matriz) && (j+l) >= 0 && (j+l) < len(matriz[i]) {
							tempSoma += matriz[i+k][j+l]
						} else {
							continue
						}
					}
				}
			}
			matriz_resultante[i][j] = tempSoma
		}
	}

	for i := 0; i < len(matriz_resultante); i++ {
		for j := 0; j < len(matriz_resultante[i]); j++ {
			f.Print(matriz_resultante[i][j], " ")
		}
		f.Print("\n")
	}

}
