package main

import f "fmt"

func main() {
	var base10 int
	var base2 []int

	f.Print("Digite o valor na base 10: ")
	f.Scan(&base10)

	for base10 > 0 {
		if base10%2 == 0 {
			base2 = append(base2, 0)
		} else {
			base2 = append(base2, 1)
		}
		base10 /= 2
	}
	f.Print("Valor na base 2: ")
	for i := len(base2) - 1; i >= 0; i-- {
		f.Printf("%d", base2[i])
	}
}
