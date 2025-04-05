package main

import "fmt"

func main() {
	var ano int

	fmt.Scan(&ano)

	if ano%4 == 0 {
		if ano%100 == 0 && ano%400 != 0 {
			fmt.Print("Nao bissexto")
		} else {
			fmt.Print("Bissexto")
		}

	} else {
		fmt.Print("Nao bissexto")
	}
}
