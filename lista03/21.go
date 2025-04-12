package main

import f "fmt"

func main() {
	var b, n, Result int

	f.Print("Digite os valores da base e expoente: ")
	f.Scan(&b, &n)
	Result = b

	for i := 0; i < n-1; i++ {
		Result *= b
	}

	f.Printf("%d^%d = %d", b, n, Result)
}
