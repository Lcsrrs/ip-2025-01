package main

import f "fmt"

func main() {
	var num int
	var achou = false
	f.Print("Digite um número inteiro positivo: ")
	f.Scan(&num)

	for i := 0; i <= num/2; i++ {
		if i*(i+1)*(i+2) == num {
			achou = true
			f.Printf("O número %d é triangular dos números %d, %d e %d", num, i, i+1, i+2)
			break
		} else if i == num/2 && achou == false {
			f.Printf("O número %d não é triagular", num)
		}
	}
}
