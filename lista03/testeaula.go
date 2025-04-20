package main

import f "fmt"

func main() {
	var array = []float64{0, 2, 4, 7, 9}

	f.Println(BuscaBinaria(array, 4))
}

func BuscaBinaria(l []float64, x float64) int {
	n := len(l)
	e := 0
	d := n - 1

	for e <= d {
		m := (e + d) / 2
		if l[m] == x {
			return m
		} else if l[m] < x {
			e = m + 1
		} else {
			d = m - 1
		}
	}
	return -1

}
