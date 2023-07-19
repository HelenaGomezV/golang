package main

import "fmt"

type Car struct {
	year  int
	brand string
}

func main() {
	myCar := Car{year: 1987, brand: "Ford"}
	fmt.Println(myCar)

}
