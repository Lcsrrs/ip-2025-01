package main

import (
	f "fmt"
)

func main() {
	var N, K, M int
	var matriz1 [][]int
	var matriz2 [][]int
	var matriz_resultante [][]int
	f.Scan(&N, &K, &M)
	matriz1 = make([][]int, N)
	matriz2 = make([][]int, K)
	matriz_resultante = make([][]int, N)
	for i := range matriz1 {
		matriz1[i] = make([]int, K)
	}
	for i := range matriz2 {
		matriz2[i] = make([]int, M)
	}
	for i := range matriz_resultante {
		matriz_resultante[i] = make([]int, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			f.Scan(&matriz1[i][j])
		}
	}
	for i := 0; i < K; i++ {
		for j := 0; j < M; j++ {
			f.Scan(&matriz2[i][j])
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			tempMult := 0
			for l := 0; l < K; l++ {
				tempMult += matriz1[i][l] * matriz2[l][j]
			}
			matriz_resultante[i][j] = tempMult
		}
	}

	for i := 0; i < len(matriz_resultante); i++ {
		for j := 0; j < len(matriz_resultante[i]); j++ {
			f.Print(matriz_resultante[i][j], " ")
		}
		f.Print("\n")
	}

}
