package main

import "fmt"

func main() {
	//defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//buena practica para cerrar conecciones, son buenas practicas utilizar un solo difer por funcion

	// CONTINUE AND BREAK

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		// continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		// continue
		if i == 8 {
			fmt.Println("Es 8")
			break
		}

	}

}
