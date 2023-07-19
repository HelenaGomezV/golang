package main

import (
	"cursoPlatzi/src/mypackages"
	"fmt"
)

func main() {
	var myCar mypackages.Car
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	// var myOtherCar mypackages.carPrivate
	// fmt.Println(myCar)

	mypackages.PrintMessage()
}
