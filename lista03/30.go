package main

import (
	f "fmt"
	m "math"
)

func main() {
	for i := 0.0; i <= 20; i += 0.5 {
		f.Printf("Volume da esfera de raio %.1f = %.2f\n", i, (4/3)*m.Pi*m.Pow(i, 3))
	}
}
