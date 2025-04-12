package main

import (
	f "fmt"
	m "math"
)

func main() {
	for i := 0.0; i <= 6.3; i += 0.1 {
		f.Printf("Sen %.2f = %.2f\n", i, i-m.Pow(i, 3)/6+m.Pow(i, 5)/120-m.Pow(i, 7)/5040)
	}
}
