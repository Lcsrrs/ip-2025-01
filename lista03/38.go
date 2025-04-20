package main

import (
	f "fmt"
	"strconv"
)

func main() {
	var somaDigitos, verificadoresCPF, verificadoresCalc, CPF9, CPF10 int
	var CPFstring string

	//Captura do CPF em string para caso ele comece com o dígito 0:
	f.Print("Digite o CPF: ")
	f.Scan(&CPFstring)

	//Verificação do tamanho do número digitado:
	if len((CPFstring)) != 11 {
		f.Print("CPF inválido")
		return
	}

	//Conversão do valor digitado para int e, em caso de erro, retornar inválido
	CPF, err := strconv.Atoi(CPFstring)
	if err != nil {
		f.Print("CPF inválido")
		return
	}

	verificadoresCPF = CPF % 100
	CPF9 = CPF / 100
	CPF10 = CPF / 10

	//Primeira soma de dígitos:
	for i := 2; i <= 10; i++ {
		somaDigitos += i * (CPF9 % 10)
		CPF9 /= 10
	}

	if somaDigitos%11 < 2 {
		verificadoresCalc = 0
	} else {
		verificadoresCalc = 10 * (11 - somaDigitos%11)
	}

	//Zerando a variável para reutilizá-la na segunda soma
	somaDigitos = 0

	//Segunda soma de dígitos
	for i := 2; i <= 11; i++ {
		somaDigitos += i * (CPF10 % 10)
		CPF10 /= 10
	}

	if somaDigitos%11 < 2 {
		verificadoresCalc += 0
	} else {
		verificadoresCalc += (11 - somaDigitos%11)
	}

	if verificadoresCPF == verificadoresCalc {
		f.Print("CPF válido")
	} else {
		f.Print("CPF inválido")
	}

}
