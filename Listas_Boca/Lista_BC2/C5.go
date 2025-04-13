package main

import f "fmt"

func main() {
	var num int
	var achou = false
	f.Scan(&num)

	for i := 0; i <= num/2; i++ {
		if i*(i+1)*(i+2) == num {
			achou = true
			f.Printf("%d eh triangular", num)
			break
		} else if i == num/2 && achou == false {
			f.Printf("%d nao eh triangular", num)
		}
	}
}
