<<<<<<< HEAD
package main

import "fmt"

func main() {
	var matricula int
	var horasExt, descontoINSS, descontoIR float64
	const salMin = 788.00
	const valorHrEx = 10.00

	fmt.Print("Digite a matrícula do funcionário e a quantidade de horas extras: ")
	fmt.Scan(&matricula, &horasExt)

	salarioHrEx := horasExt * valorHrEx
	salarioBruto := 3*salMin + salarioHrEx
	if salarioBruto > 1500 {
		descontoINSS = 0.12 * salarioBruto
	} else {
		descontoINSS = 0
	}

	if salarioBruto > 2000 {
		descontoIR = 0.2 * salarioBruto
	} else {
		descontoIR = 0
	}

	salLiq := salarioBruto - descontoINSS - descontoIR

	fmt.Printf("O salário final do funcionário %d é igual a: R$ %.2f", matricula, salLiq)
}
=======
package main

import "fmt"

func main() {
	var matricula int
	var horasExt, descontoINSS, descontoIR float64
	const salMin = 788.00
	const valorHrEx = 10.00

	fmt.Print("Digite a matrícula do funcionário e a quantidade de horas extras: ")
	fmt.Scan(&matricula, &horasExt)

	salarioHrEx := horasExt * valorHrEx
	salarioBruto := 3*salMin + salarioHrEx
	if salarioBruto > 1500 {
		descontoINSS = 0.12 * salarioBruto
	} else {
		descontoINSS = 0
	}

	if salarioBruto > 2000 {
		descontoIR = 0.2 * salarioBruto
	} else {
		descontoIR = 0
	}

	salLiq := salarioBruto - descontoINSS - descontoIR

	fmt.Printf("O salário final do funcionário %d é igual a: R$ %.2f", matricula, salLiq)
}
>>>>>>> bf15460b564febf0a3bad8ba55fdcf48a890b0ef
