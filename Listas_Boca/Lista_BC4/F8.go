package main

import (
	f "fmt"
	m "math"
)

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
		softmax := 0.0
		somaDen := 0.0
		maior := matriz[i][0]
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] > maior {
				maior = matriz[i][j]
			}
		}
		for j := 0; j < len(matriz[i]); j++ {
			somaDen += m.Pow(m.E, float64(matriz[i][j])-float64(maior))
		}
		for k := 0; k < len(matriz[i]); k++ {
			softmax = m.Pow(m.E, float64(matriz[i][k])-float64(maior)) / somaDen
			matriz_resultante[i][k] = softmax
		}
	}

	for i := 0; i < len(matriz_resultante); i++ {
		for j := 0; j < len(matriz_resultante[i]); j++ {
			f.Printf("%.6f ", matriz_resultante[i][j])
		}
		f.Print("\n")
	}
}
