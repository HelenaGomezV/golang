package main

import "fmt"

func main() {

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("impar")
	}

	value := 200

	switch {
	case value > 100:
		fmt.Println("Es mayor 100")
	case value < 0:
		fmt.Println("Es menor que cero")
	default:
		fmt.Print("No condion")

	}

}
