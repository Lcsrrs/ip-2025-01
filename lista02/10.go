<<<<<<< HEAD
package main

import "fmt"

func main() {
	var destino, tipo int

	fmt.Print("Digite o ID do destino, sendo:\n1 - Região Norte\n2 - Região Nordeste\n3 - Região Centro Oeste\n4 - Região Sul\n")
	fmt.Scan(&destino)
	var check = 0
	for i := 1; i < 5; i++ {
		if destino == i {
			check = 1
		}
	}
	if check == 0 {
		fmt.Print("Destino inválido")
		return
	}

	fmt.Print("Digite o tipo de viagem, sendo:\n 1 - Só ida\n 2 - Ida e volta\n")
	fmt.Scan(&tipo)

	if tipo != 1 && tipo != 2 {
		fmt.Print("Tipo inválido")
		return
	}
	if destino == 1 {
		if tipo == 1 {
			fmt.Print("O valor da viagem é R$ 500,00")
		} else {
			fmt.Print("O valor da viagem é R$ 900,00")
		}
	} else if destino == 2 {
		if tipo == 1 {
			fmt.Print("O valor da viagem é R$ 350,00")
		} else {
			fmt.Print("O valor da viagem é R$ 650,00")
		}
	} else if destino == 3 {
		if tipo == 1 {
			fmt.Print("O valor da viagem é R$ 350,00")
		} else {
			fmt.Print("O valor da viagem é R$ 600,00")
		}
	} else {
		if tipo == 1 {
			fmt.Print("O valor da viagem é R$ 300,00")
		} else {
			fmt.Print("O valor da viagem é R$ 550,00")
		}

	}
}
=======
package main

import "fmt"

func main() {
	var destino, tipo int

	fmt.Print("Digite o ID do destino, sendo:\n1 - Região Norte\n2 - Região Nordeste\n3 - Região Centro Oeste\n4 - Região Sul\n")
	fmt.Scan(&destino)
	var check = 0
	for i := 1; i < 5; i++ {
		if destino == i {
			check = 1
		}
	}
	if check == 0 {
		fmt.Print("Destino inválido")
		return
	}

	fmt.Print("Digite o tipo de viagem, sendo:\n 1 - Só ida\n 2 - Ida e volta\n")
	fmt.Scan(&tipo)

	if tipo != 1 && tipo != 2 {
		fmt.Print("Tipo inválido")
		return
	}
	if destino == 1 {
		if tipo == 1 {
			fmt.Print("O valor da viagem é R$ 500,00")
		} else {
			fmt.Print("O valor da viagem é R$ 900,00")
		}
	} else if destino == 2 {
		if tipo == 1 {
			fmt.Print("O valor da viagem é R$ 350,00")
		} else {
			fmt.Print("O valor da viagem é R$ 650,00")
		}
	} else if destino == 3 {
		if tipo == 1 {
			fmt.Print("O valor da viagem é R$ 350,00")
		} else {
			fmt.Print("O valor da viagem é R$ 600,00")
		}
	} else {
		if tipo == 1 {
			fmt.Print("O valor da viagem é R$ 300,00")
		} else {
			fmt.Print("O valor da viagem é R$ 550,00")
		}

	}
}
>>>>>>> bf15460b564febf0a3bad8ba55fdcf48a890b0ef
