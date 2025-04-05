<<<<<<< HEAD
package main

import "fmt"

func main() {
	var id int
	var n1, n2, n3, exercicios, media float64

	fmt.Print("Digite a identificação do aluno e suas notas 1, 2 e 3 e a média dos exercícios do aluno: ")
	fmt.Scan(&id, &n1, &n2, &n3, &exercicios)

	media = (n1 + n2*2 + n3*3 + exercicios) / 7

	fmt.Printf("Para o aluno de ID %d:\nNota 1: %.2f\nNota 2: %.2f\nNota 3: %.2f\nMédia dos exercícios: %.2f\nMédia de aproveitamento: %.2f\n", id, n1, n2, n3, exercicios, media)

	if media < 4 {
		fmt.Print("Conceito: E\nResultado: REPROVADO")
	} else if media < 6 {
		fmt.Print("Conceito: D\nResultado: REPROVADO")
	} else if media < 7.5 {
		fmt.Print("Conceito: C\nResultado: APROVADO")
	} else if media < 9 {
		fmt.Print("Conceito: B\nResultado: APROVADO")
	} else {
		fmt.Print("Conceito: A\nResultado: APROVADO")
	}
}
=======
package main

import "fmt"

func main() {
	var id int
	var n1, n2, n3, exercicios, media float64

	fmt.Print("Digite a identificação do aluno e suas notas 1, 2 e 3 e a média dos exercícios do aluno: ")
	fmt.Scan(&id, &n1, &n2, &n3, &exercicios)

	media = (n1 + n2*2 + n3*3 + exercicios) / 7

	fmt.Printf("Para o aluno de ID %d:\nNota 1: %.2f\nNota 2: %.2f\nNota 3: %.2f\nMédia dos exercícios: %.2f\nMédia de aproveitamento: %.2f\n", id, n1, n2, n3, exercicios, media)

	if media < 4 {
		fmt.Print("Conceito: E\nResultado: REPROVADO")
	} else if media < 6 {
		fmt.Print("Conceito: D\nResultado: REPROVADO")
	} else if media < 7.5 {
		fmt.Print("Conceito: C\nResultado: APROVADO")
	} else if media < 9 {
		fmt.Print("Conceito: B\nResultado: APROVADO")
	} else {
		fmt.Print("Conceito: A\nResultado: APROVADO")
	}
}
>>>>>>> bf15460b564febf0a3bad8ba55fdcf48a890b0ef
