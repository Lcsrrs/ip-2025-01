package main

import f "fmt"

func main() {
	var x, sin, i, cont, nx float64
	f.Scan(&x)
	cont = x / 0.1
	nx = 0

	for i = 0; i < cont+1; i++ {
		sin = nx - nx*nx*nx/6 + nx*nx*nx*nx*nx/120 - nx*nx*nx*nx*nx*nx*nx/5040
		f.Printf("%.1f %.4f\n", nx, sin)
		nx += 0.1
	}
}
