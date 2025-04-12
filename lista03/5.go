package main

import (
	"fmt"
	f "fmt"
)

func main() {
	var continuar = 1
	var peso, altura, mediaAltura1020, somaAltura1020 float64
	var cont50, idade, contMenos40kg, contTotal, cont1020 int

	for continuar == 1 {
		f.Print("Digite o peso, idade e altura: ")
		fmt.Scan(&peso, &idade, &altura)
		contTotal++

		if idade > 50 {
			cont50++
		}

		if idade >= 10 && idade <= 20 {
			cont1020++
			somaAltura1020 += somaAltura1020
		}

		if peso < 40 {
			contMenos40kg++
		}

		fmt.Print("Valores adicionados\nDeseja continuar? (1 - Sim, 2- Não)")
		f.Scan(&continuar)
	}

	mediaAltura1020 = somaAltura1020 / float64(cont1020)

	fmt.Printf("Pessoas com idade superior a 50: %d\nMédia altura pessoas entre 10 e 20 anos: %.2f\nPercentagem pessoas com peso inferior a 40kg: %.2f", cont50, mediaAltura1020, float64(contMenos40kg)*100/float64(contTotal))
}
