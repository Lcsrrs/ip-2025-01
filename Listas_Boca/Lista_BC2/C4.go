package main

import "fmt"

func main() {
	var num = 1
	var achou = false

	for num > 0 {
		achou = false
		fmt.Scan(&num)

		if num == 0 {
			return
		}

		for i := 0; i <= num/2; i++ {
			if i*i == num {
				achou = true
				fmt.Printf("%d eh quadrado perfeito\n", num)
				break
			} else if i == num/2 && achou == false {
				fmt.Printf("%d nao eh quadrado perfeito\n", num)
			}
		}
	}
}
