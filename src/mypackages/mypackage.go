package mypackages

import "fmt"

// Car con acceso publico
type Car struct {
	Year  int
	Brand string
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage() {
	fmt.Println("Print Hello")
}
