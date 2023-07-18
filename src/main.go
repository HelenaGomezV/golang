package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaracion Variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

	//operaciones artimeticas
	// suma
	x := 10
	y := 50
	result := x + y
	fmt.Println("Suma:", result)

	//resta
	result = x - y
	fmt.Println("Resta:", result)

	//multiplicacion
	result = x * y
	fmt.Println("Multi:", result)

	//Division
	result = x / y
	fmt.Println("Division:", result)

	//modulo
	result = x % y
	fmt.Println("Modulo:", result)

	//Incremental
	result = x + 1
	fmt.Println("Incremental:", result)

	//Incremental
	result = x - 1
	fmt.Println("Decremental:", result)

	// area de un rectangulo

	result = base * altura / 2
	fmt.Println("area Rectangulo:", result)

}
