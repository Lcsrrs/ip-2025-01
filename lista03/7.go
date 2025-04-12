package main

import f "fmt"

func main() {
	var num, soma, media, maior, menor, somaPar, mediaPar float64
	var quantidade, ePar, eImpar int

	for num != 30000 {
		f.Print("Digite o número: ")
		f.Scan(&num)

		if num == 30000 {
			break
		}

		soma += num
		quantidade++
		if num > maior || quantidade == 1 {
			maior = num
		}
		if num < menor || quantidade == 1 {
			menor = num
		}
		if int(num)%2 == 0 {
			somaPar += num
			ePar++
		} else {
			eImpar++
		}
	}
	media = soma / float64(quantidade)
	mediaPar = somaPar / float64(ePar)

	f.Printf("a) Soma dos números digitados: %.2f\nb) Quantidade de números digitados: %d\nc) Média dos números digitados: %.2f\nd) Maior número digitado: %.2f\ne) Menor número digitado: %.2f\nf) Média dos números pares: %.2f\ng) Percentagem de números ímpares: %.2f", soma, quantidade, media, maior, menor, mediaPar, float64(eImpar)*100/float64(quantidade))
}
