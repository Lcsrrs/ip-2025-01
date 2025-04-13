package main

import f "fmt"

func main() {
	var l, r, numeros, soma int
	f.Scan(&l, &r)

	if l < 1 || r > 10000 || l > r {
		return
	}

	if l == r {
		if l%2 == 0 {
			f.Print(l)
			f.Print("\n", l)
			return
		} else {
			f.Print("0")
			f.Print("\n", "0")
			return
		}
	}

	for i := l; i <= r; i++ {
		if i%2 == 0 {
			numeros++
			soma += i
		}
	}
	f.Print(soma)
	f.Print("\n", soma/numeros)
}
