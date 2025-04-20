package main

import f "fmt"

func main() {
	var n1, n2 float64
	var aprovado, exame, reprovado, N int

	f.Scan(&N)

	for i := 1; i <= N; i++ {
		f.Scan(&n1, &n2)
		if (n1+n2)/2 <= 3 {
			f.Printf("Aluno %d: Reprovado\n", i)
			reprovado++
		} else if (n1+n2)/2 < 7 {
			f.Printf("Aluno %d: Exame\n", i)
			exame++
		} else {
			f.Printf("Aluno %d: Aprovado\n", i)
			aprovado++
		}
	}

	f.Printf("Total Aprovados: %d\n", aprovado)
	f.Printf("Total Exame: %d\n", exame)
	f.Printf("Total Reprovados: %d\n", reprovado)

}
