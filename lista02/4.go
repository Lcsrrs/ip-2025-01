<<<<<<< HEAD
package main

import (
	"fmt"
	"math"
)

func main() {
	var x1 float64

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x1)

	if x1 >= 0 {
		fmt.Print("A raiz quadrada do número positivo é: ", math.Sqrt(x1))
	} else {
		fmt.Print("O quadrado do número negativo é: ", x1*x1)
	}
}
=======
package main

import (
	"fmt"
	"math"
)

func main() {
	var x1 float64

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x1)

	if x1 >= 0 {
		fmt.Print("A raiz quadrada do número positivo é: ", math.Sqrt(x1))
	} else {
		fmt.Print("O quadrado do número negativo é: ", x1*x1)
	}
}
>>>>>>> bf15460b564febf0a3bad8ba55fdcf48a890b0ef
