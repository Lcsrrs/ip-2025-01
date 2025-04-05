package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	var faltas int

	fmt.Scan(&n1, &n2, &n3, &faltas)

	if faltas > 16 {
		fmt.Print("Reprovado por falta")
	} else {
		var media float64

		media = (n1 + n2 + n3) / 3

		if media < 5 {
			fmt.Print("Reprovado")
		} else if media < 7 {
			fmt.Print("Prova final")
		} else {
			fmt.Print("Aprovado")
		}
	}
}
