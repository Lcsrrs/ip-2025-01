package main

import f "fmt"

func main() {
	var linhas, colunas, tempSoma int
	var matriz [][]int
	var matrizF [][]int
	var matriz_resultante [][]int

	f.Scan(&linhas, &colunas)

	matriz = make([][]int, linhas)
	matrizF = make([][]int, 3)
	matriz_resultante = make([][]int, linhas-2)

	for i := range matriz {
		matriz[i] = make([]int, colunas)
	}
	for i := range matrizF {
		matrizF[i] = make([]int, 3)
	}
	for i := range matriz_resultante {
		matriz_resultante[i] = make([]int, colunas-2)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			f.Scan(&matriz[i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			f.Scan(&matrizF[i][j])
		}
	}

	for i := 0; i < len(matriz_resultante); i++ {
		for j := 0; j < len(matriz_resultante[i]); j++ {
			tempSoma = 0
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					tempSoma += matriz[i+di+1][j+dj+1] * matrizF[di+1][dj+1]
				}
			}
			//f.Print("temp soma: ", tempSoma, " ")
			matriz_resultante[i][j] = tempSoma
		}
	}

	for i := 0; i < len(matriz_resultante); i++ {
		for j := 0; j < len(matriz_resultante[i]); j++ {
			f.Printf("%d ", matriz_resultante[i][j])
		}
		f.Print("\n")
	}
}
