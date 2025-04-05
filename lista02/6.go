<<<<<<< HEAD
package main

import "fmt"

func main() {
	var x1, x2 int

	fmt.Print("Digite dois números: ")
	fmt.Scan(&x1, &x2)

	if x1%x2 == 0 {
		fmt.Print("O número ", x1, " é divisível pelo número ", x2)
	} else {
		fmt.Print("O número ", x1, " não é divisível pelo número ", x2)
	}
}
=======
package main

import "fmt"

func main() {
	var x1, x2 int

	fmt.Print("Digite dois números: ")
	fmt.Scan(&x1, &x2)

	if x1%x2 == 0 {
		fmt.Print("O número ", x1, " é divisível pelo número ", x2)
	} else {
		fmt.Print("O número ", x1, " não é divisível pelo número ", x2)
	}
}
>>>>>>> bf15460b564febf0a3bad8ba55fdcf48a890b0ef
