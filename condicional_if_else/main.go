package main

import "fmt"

func main() {

	var n int

	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println(n, "ES NEUTRO")
	} else if n%2 == 0 {
		fmt.Println(n, "ES PAR")
	} else {
		fmt.Println(n, "ES IMPAR")
	}
}
