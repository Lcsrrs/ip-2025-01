package main

import "fmt"

func main() {
	var idade int

	fmt.Scan(&idade)

	if idade >= 5 && idade <= 7 {
		fmt.Print("Infantil A")
	} else if idade <= 10 {
		fmt.Print("Infantil B")
	} else if idade <= 13 {
		fmt.Print("Juvenil A")
	} else if idade <= 17 {
		fmt.Print("Juvenil B")
	} else {
		fmt.Print("Adulto")
	}
}
