package main

import f "fmt"

func main() {
	var n1, n2, somaClasse float64
	var aprovado, exame, reprovado, N int

	f.Scan(&N)

	for i := 0; i < N; i++ {
		f.Scan(&n1, &n2)
		somaClasse = somaClasse + n1 + n2
		if (n1+n2)/2 < 3 {
			f.Printf("Aluno %d: Reprovado\n", i+1)
			reprovado++
		} else if (n1+n2)/2 < 7 {
			f.Printf("Aluno %d: Exame\n", i+1)
			exame++
		} else {
			f.Printf("Aluno %d: Aprovado\n", i+1)
			aprovado++
		}
	}

	f.Printf("Total Aprovados: %d\nTotal Exame: %d\nTotal Reprovados: %d", aprovado, exame, reprovado)

}
