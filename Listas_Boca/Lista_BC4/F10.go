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

	cima := 0
	baixo := len(matriz) - 1
	esquerda := 0
	direita := len(matriz[0]) - 1

	for cima <= baixo && esquerda <= direita {
		for i := esquerda; i <= direita; i++ {
			f.Print(matriz[cima][i], " ")
		}
		cima++
		for i := cima; i <= baixo; i++ {
			f.Print(matriz[i][direita], " ")
		}
		direita--
		if cima <= baixo {
			for i := direita; i >= esquerda; i-- {
				f.Print(matriz[baixo][i], " ")
			}
		}
		baixo--
		if esquerda <= direita {
			for i := baixo; i >= cima; i-- {
				f.Print(matriz[i][esquerda], " ")
			}
		}
		esquerda++
	}
}
