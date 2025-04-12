package main

import "fmt"

func main() {
	var num = 1
	var achou = false

	for num > 0 {
		achou = false
		fmt.Print("\nDigite o número: ")
		fmt.Scan(&num)

		for i := 0; i <= num/2; i++ {
			if i*i == num {
				achou = true
				fmt.Printf("O número %d é quadrado perfeito do número %d\n", num, i)
				break
			} else if i == num/2 && achou == false {
				fmt.Print("Não é quadrado perfeito\n")
			}
		}
	}
}
