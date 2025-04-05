package main

import "fmt"

func main() {
	var x, y, z int

	fmt.Scan(&x, &y, &z)

	var maior = x
	var menor1 = y
	var menor2 = z

	if x > y && x > z {
		maior = x
		menor1 = y
		menor2 = z
	} else if y > x && y > z {
		maior = y
		menor1 = x
		menor2 = z
	} else if z > x && z > y {
		maior = z
		menor1 = x
		menor2 = y
	}

	if menor1+menor2 <= maior {
		fmt.Print("Nao forma triangulo")
	} else if menor1 == menor2 && menor2 == maior {
		fmt.Print("Equilatero")
	} else if menor1 == menor2 || menor1 == maior || menor2 == maior {
		fmt.Print("Isosceles")
	} else {
		fmt.Print("Escaleno")
	}
}
