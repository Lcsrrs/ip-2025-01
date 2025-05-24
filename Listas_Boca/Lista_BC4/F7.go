package main

import f "fmt"

func main() {
	var linhas, colunas int
	var matriz [][]int
	var matriz_resultante [][]float64
	f.Scan(&linhas, &colunas)
	matriz = make([][]int, linhas)
	matriz_resultante = make([][]float64, linhas)
	for i := range matriz {
		matriz[i] = make([]int, colunas)
		matriz_resultante[i] = make([]float64, colunas)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			f.Scan(&matriz[i][j])
		}
	}

	for i := 0; i < len(matriz); i++ {
		maior := matriz[i][0]
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] > maior {
				maior = matriz[i][j]
			}
		}

		for j := 0; j < len(matriz[i]); j++ {
			matriz_resultante[i][j] = float64(matriz[i][j]) / float64(maior)
		}
	}

	for i := 0; i < len(matriz_resultante); i++ {
		for j := 0; j < len(matriz_resultante[i]); j++ {
			f.Printf("%.6f ", matriz_resultante[i][j])
		}
		f.Print("\n")
	}
}
