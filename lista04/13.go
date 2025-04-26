package main

import f "fmt"

func main() {
	type funcionario struct {
		numeroEmpregado int
		mesesEmpregado  int
	}
	var empregados []funcionario
	var temp, menor funcionario
	var numeroFunc, indMenor, cont int

	for numeroFunc < 100 {
		f.Printf("Digite a matrícula e a quantidade de meses do %dº empregado: ", numeroFunc+1)
		f.Scan(&temp.numeroEmpregado, &temp.mesesEmpregado)
		if temp.mesesEmpregado == 0 && temp.numeroEmpregado == 0 {
			break
		} else {
			empregados = append(empregados, temp)
			numeroFunc++
		}
	}

	menor.mesesEmpregado = 0
	menor.numeroEmpregado = 0

	//Ordenar por tempo de casa:
	for i := 0; i < len(empregados); i++ {
		menor.mesesEmpregado = empregados[i].mesesEmpregado
		indMenor = i
		for j := i + 1; j < len(empregados); j++ {
			if empregados[j].mesesEmpregado < menor.mesesEmpregado {
				menor.mesesEmpregado = empregados[j].mesesEmpregado
				indMenor = j
			}
		}
		temp = empregados[i]
		empregados[i] = empregados[indMenor]
		empregados[indMenor] = temp
		cont++
		if cont == 3 {
			break
		}
	}

	cont = 0
	f.Print("Os três funcionários mais novos são:\n")
	for i := 0; i < len(empregados); i++ {
		f.Printf("Funcionário %d, com %d meses de casa\n", empregados[i].numeroEmpregado, empregados[i].mesesEmpregado)
		cont++
		if cont == 3 {
			break
		}
	}

}
