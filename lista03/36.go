package main

import (
	f "fmt"
	"strconv"
)

func main() {
	var base10 int
	var base16 []string

	f.Print("Digite o valor na base 10: ")
	f.Scan(&base10)

	for base10 > 0 {
		switch {
		case base10%16 < 10:
			base16 = append(base16, strconv.Itoa(base10%16))
		case base10%16 == 10:
			base16 = append(base16, "A")
		case base10%16 == 11:
			base16 = append(base16, "B")
		case base10%16 == 12:
			base16 = append(base16, "C")
		case base10%16 == 13:
			base16 = append(base16, "D")
		case base10%16 == 14:
			base16 = append(base16, "E")
		case base10%16 == 15:
			base16 = append(base16, "F")
		}
		base10 /= 16
	}

	f.Print("Valor na base 16: ")
	for i := len(base16) - 1; i >= 0; i-- {
		f.Printf("%s", base16[i])
	}
}
