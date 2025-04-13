package main

import f "fmt"

func main() {
	var base, expoente int
	potencia := 1
	f.Scan(&base, &expoente)

	if expoente == 0 {
		f.Print("1")
		return
	}

	for i := 1; i <= expoente; i++ {
		potencia *= base
	}

	f.Print(potencia)
}
