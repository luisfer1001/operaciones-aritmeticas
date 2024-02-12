package main

import "fmt"

func main() {
	var (
		n      int
		salida string
	)

	fmt.Print("Ingrese un n uemero de 1 a 5: ")
	fmt.Scan(&n)

	switch n {
	case 1:
		salida = "UNO"
	case 2:
		salida = "DOS"
	case 3:
		salida = "TRES"
	case 4:
		salida = "CUATRO"
	case 5:
		salida = "CINCO"
	default:
		salida = "NO ESTA ENTRE 1 Y 5"
	}

	//SALIDA
	fmt.Println(n, salida)

}
