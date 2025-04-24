package main

import f "fmt"

func main() {
	var ano int

	f.Scan(&ano)

	anoBase := ano - 2004

	for anoBase > 12 {
		anoBase -= 12
	}

	switch anoBase {
	case 0:
		f.Print("Macaco")
	case 1:
		f.Print("Galo")
	case 2:
		f.Print("Cao")
	case 3:
		f.Print("Porco")
	case 4:
		f.Print("Rato")
	case 5:
		f.Print("Boi")
	case 6:
		f.Print("Tigre")
	case 7:
		f.Print("Coelho")
	case 8:
		f.Print("Dragao")
	case 9:
		f.Print("Serpente")
	case 10:
		f.Print("Cavalo")
	case 11:
		f.Print("Cabra")
	}
}
