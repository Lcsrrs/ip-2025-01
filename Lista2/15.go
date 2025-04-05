package main

import "fmt"

func main() {
	var dia, mes, ano int

	fmt.Print("DIgite dia, mês e ano separados por espaço: ")
	fmt.Scan(&dia, &mes, &ano)

	if mes == 1 {
		fmt.Printf("Data: %d de janeiro de %d", dia, ano)
	}
	if mes == 2 {
		fmt.Printf("Data: %d de fevereiro de %d", dia, ano)
	}
	if mes == 3 {
		fmt.Printf("Data: %d de março de %d", dia, ano)
	}
	if mes == 4 {
		fmt.Printf("Data: %d de abril de %d", dia, ano)
	}
	if mes == 5 {
		fmt.Printf("Data: %d de maio de %d", dia, ano)
	}
	if mes == 6 {
		fmt.Printf("Data: %d de junho de %d", dia, ano)
	}
	if mes == 7 {
		fmt.Printf("Data: %d de julho de %d", dia, ano)
	}
	if mes == 8 {
		fmt.Printf("Data: %d de agosto de %d", dia, ano)
	}
	if mes == 9 {
		fmt.Printf("Data: %d de setembro de %d", dia, ano)
	}
	if mes == 10 {
		fmt.Printf("Data: %d de outubro de %d", dia, ano)
	}
	if mes == 11 {
		fmt.Printf("Data: %d de novembro de %d", dia, ano)
	}
	if mes == 12 {
		fmt.Printf("Data: %d de dezembro de %d", dia, ano)
	}

}
