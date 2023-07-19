package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("True")
	}

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("True")
	}

	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

}
