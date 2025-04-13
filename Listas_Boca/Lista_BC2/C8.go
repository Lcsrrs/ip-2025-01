package main

import (
	f "fmt"
	m "math"
)

func main() {
	var x, sin float64
	f.Scan(&x)

	for i := 0.0; i <= x; i += 0.1 {
		sin = i - m.Pow(i, 3)/6 + m.Pow(i, 5)/120 - m.Pow(i, 7)/5040
		f.Printf("%f %f\n", i, sin)
	}
}
