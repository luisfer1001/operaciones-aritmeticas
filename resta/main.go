package main

import "fmt"

func main() {
	var a, b int

	//Entrada
	fmt.Print("Ingreso el valor de a: ")
	fmt.Scanln(&a)

	fmt.Print("Ingreso el valor de b: ")
	fmt.Scanln(&b)

	//Proceso
	suma := a + b
	resta := a - b
	multi := a * b
	divi := a / b
	mod := a % b

	//Igualdad ==
	fmt.Println("¿ Son igales ? => ", a == b)
	//Distintos !=
	fmt.Println(" ¿ Son distintos ? => ", a != b)
	//Menor que <
	fmt.Println(" ¿ El primero es menor que el segundo ? => ", a < b)
	//Menor o igual que <=
	fmt.Println(" ¿ El primero es menor o igual que el segundo ? => ", a <= b)
	//Mayor que >
	fmt.Println(" ¿ El primero es mayor que el segundo ? => ", a > b)
	//Mayor o igual que >=
	fmt.Println(" ¿ El primero es mayor o igual que el segundo ? => ", a >= b)

	//Salida
	fmt.Println("La Suma es: ", suma)
	fmt.Println("La Resta es: ", resta)
	fmt.Println("La Multiplicacion es: ", multi)
	fmt.Println("La Division es: ", divi)
	fmt.Println("El Modulo suma es: ", mod)
}
