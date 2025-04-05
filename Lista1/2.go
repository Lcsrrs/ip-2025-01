package main

import "fmt"

func main() {
	var (
		publico, popular, geral, arquibancada, cadeiras float32
	)
	var casos int
	var todos_casos []float32
	fmt.Print("Quantidade de casos: ")
	fmt.Scan(&casos)
	for i := 0; i < casos; i++ {
		fmt.Print("Público total, popular, geral, arquibancada e cadeiras: ")
		fmt.Scan(&publico, &popular, &geral, &arquibancada, &cadeiras)
		todos_casos = append(todos_casos, publico)
		todos_casos = append(todos_casos, popular)
		todos_casos = append(todos_casos, geral)
		todos_casos = append(todos_casos, arquibancada)
		todos_casos = append(todos_casos, cadeiras)
	}
	var renda float32
	var cont_casos = 1
	for i := 0; i < casos+4*casos; i = i + 5 {
		renda = todos_casos[i]*todos_casos[i+1]/100 + todos_casos[i]*5*todos_casos[i+2] + todos_casos[i]*10*todos_casos[i+3] + todos_casos[i]*20*todos_casos[i+4]
		fmt.Print("A renda do jogo N, ", cont_casos, " é: ", renda, "\n")
		cont_casos++
	}
}
