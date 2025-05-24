package main

import f "fmt"

func main() {
	var linhas, colunas, tempSoma int
	var matriz [][]int
	var somaLinha []int
	var somaColuna []int
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

	// Soma linha:
	for i := 0; i < len(matriz); i++ {
		tempSoma = 0
		for j := 0; j < len(matriz[i]); j++ {
			tempSoma += matriz[i][j]
		}
		somaLinha = append(somaLinha, tempSoma)
	}

	// Soma coluna:
	for i := 0; i < colunas; i++ {
		tempSoma = 0
		for j := 0; j < linhas; j++ {
			tempSoma += matriz[j][i]
		}
		somaColuna = append(somaColuna, tempSoma)
	}

	//Imprimindo soma das linhas:
	for i := 0; i < len(somaLinha); i++ {
		f.Print(somaLinha[i], " ")
	}

	f.Print("\n")

	//Imprimindo soma das colunas:
	for i := 0; i < len(somaColuna); i++ {
		f.Print(somaColuna[i], " ")
	}
}
