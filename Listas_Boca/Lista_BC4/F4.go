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

	// Verificação linha:
	crescenteLinha := true
	for i := 0; i < len(matriz); i++ {
		menor := matriz[i][0]
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] >= menor {
				menor = matriz[i][j]
			} else {
				crescenteLinha = false
				break
			}
		}
		if crescenteLinha == false {
			break
		}
	}

	//Verificação coluna:
	crescenteColuna := true
	if crescenteLinha {
		for i := 0; i < colunas; i++ {
			menor := matriz[0][i]
			for j := 0; j < linhas; j++ {
				if matriz[j][i] >= menor {
					menor = matriz[j][i]
				} else {
					crescenteColuna = false
					break
				}
			}
			if crescenteColuna == false {
				break
			}
		}
	}

	if crescenteColuna && crescenteLinha {
		f.Print("SIM")
	} else {
		f.Print("NAO")

	}

}
