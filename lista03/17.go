package main

import f "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			f.Print(i+1, "x", j+1, " ")
		}
		f.Print("\n")
	}
}
