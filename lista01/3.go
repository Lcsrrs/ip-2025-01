package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int
	fmt.Print("Digite três números separados por espaço: ")
	fmt.Scan(&n1, &n2, &n3)
	//if n1 > 9 || n1 < 0 || unicode.IsLetter(n1) {
	//	fmt.Print("O primeiro número é INVÁLIDO")
	//}
	var concatenado = n1*100 + n2*10 + n3
	fmt.Print("Números concatenados: ", concatenado)
	fmt.Print("\n")
	fmt.Print("Quadrado do concatenado: ", concatenado*concatenado)
}
