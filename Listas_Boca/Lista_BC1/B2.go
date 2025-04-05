package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Print(x, " Par")
	} else {
		fmt.Print(x, " Impar")

	}
}
