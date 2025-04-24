package main

import f "fmt"

func main() {
	var n int

	f.Scan(&n)

	if n < 3 {
		f.Print("E")
	} else if n < 5 {
		f.Print("D")
	} else if n < 7 {
		f.Print("C")
	} else if n < 9 {
		f.Print("B")
	} else {
		f.Print("A")
	}
}
