package main

import f "fmt"

func main() {
	var n1, n2, somaClasse float64
	var aprovado, exame, reprovado, N int

	f.Print("Digite a quantidade N de alunos: ")
	f.Scan(&N)

	for i := 0; i < N; i++ {
		f.Print("\nDigite as duas notas: ")
		f.Scan(&n1, &n2)
		somaClasse = somaClasse + n1 + n2
		if (n1+n2)/2 < 3 {
			f.Print("Média: ", (n1+n2)/2, "\nReprovado")
			reprovado++
		} else if (n1+n2)/2 < 7 {
			f.Print("Média: ", (n1+n2)/2, "\nExame")
			exame++
		} else {
			f.Print("Média: ", (n1+n2)/2, "\nAprovado")
			aprovado++
		}
	}

	f.Printf("\nAlunos aprovados: %d\nAlunos Exame: %d\nAlunos Reprovados: %d\nMédia Classe: %.2f", aprovado, exame, reprovado, somaClasse/(2*float64(N)))

}
