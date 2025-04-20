package main

import f "fmt"

func main() {
	var pInit, D, deltaP, pMin, pAtual, Lucro, lucroMax, pMax float64
	var qInit, deltaQ, qAtual, qMax int

	f.Print("Digite o preço inicial, a quantidade de ingressos vendidos, as despesas, a variação de preço, a variação do preço, a variação de ingressos e o preço mínimo tolerado: ")
	f.Scan(&pInit, &qInit, &D, &deltaP, &deltaQ, &pMin)
	pAtual = pInit
	qAtual = qInit
	f.Print("Preco Ingressos Vendidos Lucro\n")

	for pAtual >= pMin {
		Lucro = (pAtual * float64(qAtual)) - D
		f.Printf("%.2f %d %.2f\n", pAtual, qAtual, Lucro)
		if Lucro > lucroMax {
			lucroMax = Lucro
			pMax = pAtual
			qMax = qAtual
		}
		pAtual -= deltaP
		qAtual += deltaQ
	}
	f.Printf("Lucro maximo: %.2f\nNa faixa de preco: %.2f com %d ingressos.", lucroMax, pMax, qMax)
}
