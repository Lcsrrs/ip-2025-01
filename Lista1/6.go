package main

import "fmt"

func main() {
	var (
		temperaturaC, temperaturaF float32
	)
	var casos int
	var todos_casos []float32
	fmt.Print("Quantidade de casos: ")
	fmt.Scan(&casos)
	for i := 0; i < casos; i++ {
		fmt.Print("Temperatura em Fahrenheit, caso ", i+1, " : ")
		fmt.Scan(&temperaturaF)
		todos_casos = append(todos_casos, temperaturaF)
	}
	for i := 0; i < casos; i++ {
		temperaturaC = 5 * (todos_casos[i] - 32) / 9
		fmt.Print(temperaturaF, " EQUIVALE A ", temperaturaC, " CELSIUS\n")
	}

}
