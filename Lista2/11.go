package main

import "fmt"

func main() {

	var x float32

	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)

	if x == 2 {
		fmt.Print("Resultado inválido")
	} else {
		fmt.Printf("O valor da função f(x)=8/(2-x) é igual a %f", 8/(2-x))
	}

}
