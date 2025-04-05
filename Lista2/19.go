package main

import (
	"fmt"
	"math"
)

func main() {
	var tipo int
	const pi = 3.14159

	fmt.Print("Qual tipo de sólido deseja calcular o volume?\n1 - Cone\n2 - Cilindro Reto\n3 - Esfera\n")
	fmt.Scan(&tipo)

	if tipo != 1 && tipo != 2 && tipo != 3 {
		fmt.Print("Tipo inválido")
		return
	}

	if tipo == 1 {
		var raio1, altura1, volume1, area1 float64

		fmt.Print("Digite o raio e a altura do cone: ")
		fmt.Scan(&raio1, &altura1)

		volume1 = pi * raio1 * raio1 * altura1 / 3
		area1 = pi * raio1 * math.Sqrt(raio1*raio1+altura1*altura1)
		fmt.Printf("A área do cone é igual a %.2f e o seu volume é igual a %.2f", area1, volume1)
	} else if tipo == 2 {
		var raio2, altura2, volume2, area2 float64

		fmt.Print("Digite o raio e a altura do cilindro: ")
		fmt.Scan(&raio2, &altura2)

		volume2 = pi * raio2 * raio2 * altura2
		area2 = 2 * pi * raio2 * altura2
		fmt.Printf("A área do cilindro é igual a %.2f e o seu volume é igual a %.2f", area2, volume2)
	} else {
		var raio3, volume3, area3 float64

		fmt.Print("Digite o raio da esfera: ")
		fmt.Scan(&raio3)

		area3 = 4 * pi * raio3 * raio3
		volume3 = (4 / 3) * pi * math.Pow(raio3, 3)
		fmt.Printf("A área da esfera é igual a %.2f e o seu volume é igual a %.2f", area3, volume3)
	}
}
